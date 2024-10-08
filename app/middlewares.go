package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/o1egl/paseto"
)

type Middleware func(http.Handler) http.Handler

func createStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

type WrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedWriter := &WrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusAccepted,
		}
		next.ServeHTTP(wrappedWriter, r)
		log.Println(wrappedWriter.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			writeJSON(w, http.StatusUnauthorized, M{"error": "Authorization header missing"})
			return
		}

		// Remove "Bearer " prefix if present else raise error
		bearerPrefix := "Bearer "
		if len(tokenString) > len(bearerPrefix) && strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = tokenString[len(bearerPrefix):]
		} else {
			writeJSON(w, http.StatusUnauthorized, M{"error": "Invalid token"})
			return
		}

		var jsonToken paseto.JSONToken
		var footer string
		pasetoToken := paseto.NewV2()

		if err := pasetoToken.Decrypt(tokenString, pasetoKey, &jsonToken, &footer); err != nil {
			writeJSON(w, http.StatusUnauthorized, M{"error": "Invalid token"})
			return
		}

		// Check token expiration
		if jsonToken.Expiration.Before(time.Now()) {
			http.Redirect(w, r, "/auth/login/", http.StatusUnauthorized)
			return
		}
		user_id_string := jsonToken.Subject
		user_id, err := strconv.Atoi(user_id_string)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, M{"error": err.Error()})
			return
		}
		// Fetch the user from the database
		client := GetClient()

		user, err := client.User.Get(context.Background(), user_id)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, M{"error": "User Not Found"})
			return
		}

		// Attach the user to the request context
		ctx := context.WithValue(r.Context(), userContextKey, user)

		// Pass the request to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var allowedOrigins = []string{
	"http://localhost",
	"http://127.0.0.1",
	"https://localhost",
	"https://127.0.0.1",
}

// CORS middleware to handle CORS and set headers
// We are allowing only local requests currently
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		// Check if the origin is in the list of allowed origins
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if strings.HasPrefix(origin, allowedOrigin) {
				allowed = true
				break
			}
		}

		if allowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			// Deny the request if the origin is not allowed
			writeJSON(w, http.StatusForbidden, M{"error": "Cross Origin Forbidden"})
			return
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONS Header is used just to give the headers so returning it here only
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
