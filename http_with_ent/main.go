package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Load the Env file
	if err := LoadEnv(); err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	client := GetClient()
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// enable debugging
	// client = client.Debug()

	log.Println("Connected to db")

	// Ordering by most specific wins, instead of top down approach
	// panics if both routes are equally as specific

	pasetoKey = getPasetoKey()

	user_router := http.NewServeMux()
	user_router.HandleFunc("GET /", getUsers)
	user_router.HandleFunc("GET /{id}", getUserById)
	user_router.HandleFunc("PATCH /{id}", updateUserById)
	user_router.HandleFunc("DELETE /{id}", deleteUserById)

	friends_router := http.NewServeMux()
	friends_router.HandleFunc("POST /", addFriendById)
	friends_router.HandleFunc("DELETE /", deleteFriendById)

	blog_router := http.NewServeMux()
	blog_router.HandleFunc("GET /", getBlogs)
	blog_router.HandleFunc("GET /{id}", getBlogById)
	blog_router.HandleFunc("POST /", createBlog)
	blog_router.HandleFunc("PATCH /{id}", updateBlogById)
	blog_router.HandleFunc("DELETE /{id}", deleteByBlogId)

	tags_router := http.NewServeMux()
	tags_router.HandleFunc("PATCH /{id}", updateTagById)
	tags_router.HandleFunc("GET /", getTags)

	api_router := http.NewServeMux()
	api_router.Handle("/user/", http.StripPrefix("/user", user_router))
	api_router.Handle("/blog/", http.StripPrefix("/blog", blog_router))
	api_router.Handle("/friend/", http.StripPrefix("/friend", friends_router))
	api_router.Handle("/tag/", http.StripPrefix("/tag", tags_router))

	login_router := http.NewServeMux()
	login_router.HandleFunc("POST /signout/", signOutHandler)
	login_router.HandleFunc("POST /login/", loginHandler)
	login_router.HandleFunc("POST /signup/", signUpHandler)

	router := http.NewServeMux()
	router.Handle("/auth/", http.StripPrefix("/auth", login_router))
	router.Handle("/api/", http.StripPrefix("/api", authenticateUser(api_router)))

	stack := createStack(
		logging,
		CORSMiddleware,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	// We are adding the server to a goroutine now and
	// Using a channel of kill signal to handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on port 8080")
		if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client.Close()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
