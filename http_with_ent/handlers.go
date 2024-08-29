package main

import (
	"context"
	"errors"
	"example/hello/http_with_ent/ent"
	"example/hello/http_with_ent/ent/blog"
	"example/hello/http_with_ent/ent/tag"
	"example/hello/http_with_ent/ent/user"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserDetails struct {
	Name     *string `json:"name"`
	Age      *int    `json:"age"`
	IsActive *bool   `json:"is_active"`
}

type FriendRequest struct {
	FriendID int `json:"friend_id"`
}

type BlogDetails struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Episode     *int     `json:"episode"`
	UserId      *int     `json:"user_id"`
	TagNames    []string `json:"tags"`
}

type TagUpdateRequest struct {
	Name     *string `json:"name"`
	Type     *string `json:"type"`
	Category *string `json:"category"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	users, err := client.User.Query().WithBlogs().WithFriends().All(context.Background())
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, users)
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	// Extract query parameters
	tagCategory := r.URL.Query().Get("category")

	// Build the query
	query := client.Blog.Query().WithUser().WithTags()

	if tagCategory != "" {
		query = query.Where(blog.HasTagsWith(tag.CategoryEQ(tag.Category(tagCategory))))
	}

	blogs, err := query.All(context.Background())
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, blogs)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"message": "Invalid Id received"})
		return
	}
	user, err := client.User.
		Query().
		Where(user.ID(id)).
		WithBlogs().
		WithFriends().
		Only(r.Context())

	if err != nil {
		message := ""
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// Handle the case where the entity was not found
			message = ("User with Id: " + id_string + " was not found")
		} else {
			// Handle other errors
			message = err.Error()
		}
		writeJSON(w, http.StatusBadRequest, M{"error": message})
		return
	}
	writeJSON(w, http.StatusOK, user)
}
func getBlogById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"message": "Invalid Id received"})
		return
	}
	users, err := client.Blog.Get(context.Background(), id)
	if err != nil {
		message := ""
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// Handle the case where the entity was not found
			message = ("Blog with Id: " + id_string + " was not found")
		} else {
			// Handle other errors
			message = err.Error()
		}
		writeJSON(w, http.StatusBadRequest, M{"error": message})
		return
	}
	writeJSON(w, http.StatusOK, users)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	login_json := LoginRequest{}

	if err := readJSON(w, r, &login_json); err != nil {
		writeJSON(w, http.StatusUnauthorized, M{"error": err.Error()})
		return
	}
	log.Printf("login_json: %v\n", login_json)

	user, err := client.User.
		Query().
		Where(user.Name(login_json.Name)).
		Only(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, M{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_json.Password)); err != nil {
		writeJSON(w, http.StatusUnauthorized, M{"error": "Invalid credentials"})
		return
	}

	// Create PASETO token
	token := paseto.NewV2()
	jsonToken := paseto.JSONToken{
		Subject:    strconv.Itoa(user.ID),
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(24 * time.Hour),
	}
	log.Printf("pasetoKey: %v\n", pasetoKey)
	encryptedToken, err := token.Encrypt(pasetoKey, jsonToken, nil)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, M{"error": "Internal Server Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	writeJSON(w, http.StatusOK, M{"token": encryptedToken})
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	signup_json := LoginRequest{}

	if err := readJSON(w, r, &signup_json); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}
	log.Printf("signup_json: %v\n", signup_json)

	// Check if user already exists
	_, err := client.User.
		Query().
		Where(user.Name(signup_json.Name)).
		Only(context.Background())

	if err == nil {
		writeJSON(w, http.StatusConflict, M{"error": "User already exists"})
		return
	} else if !ent.IsNotFound(err) {
		writeJSON(w, http.StatusInternalServerError, M{"error": "Internal server error"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup_json.Password), bcrypt.DefaultCost)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, M{"error": "Internal server error"})
		return
	}

	// Create the user in the database
	newUser, err := client.User.
		Create().
		SetName(signup_json.Name).
		SetPassword(string(hashedPassword)).
		Save(context.Background())

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, M{"error": "Failed to create user"})
		return
	}

	writeJSON(w, http.StatusCreated, M{"message": "User created successfully", "user_id": newUser.ID})
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	// Invalidate the token by removing it from the client's cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Set expiration to past to remove the cookie
		HttpOnly: true,
	})

	// Send a response confirming sign-out
	writeJSON(w, http.StatusOK, M{"message": "Signed out successfully"})
}

func createBlog(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	blog_json := BlogDetails{}

	if err := readJSON(w, r, &blog_json); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}
	log.Printf("blog_json: %v\n", blog_json)

	user := GetUserFromContext(r.Context())

	// Create or find tags
	var tags []*ent.Tag
	for _, tagName := range blog_json.TagNames {
		tag, err := client.Tag.Query().Where(tag.Name(tagName)).Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				// Create new tag if not found
				tag, err = client.Tag.Create().SetName(tagName).Save(context.Background())
				if err != nil {
					writeJSON(w, http.StatusInternalServerError, M{"error": "Failed to create tag"})
					return
				}
			} else {
				writeJSON(w, http.StatusInternalServerError, M{"error": "Failed to query tag"})
				return
			}
		}
		tags = append(tags, tag)
	}

	save := client.Blog.Create().SetTitle(*blog_json.Title).SetDescription(*blog_json.Description).SetUser(user)
	if blog_json.Episode != nil {
		save = save.SetEpisode(*blog_json.Episode)
	}
	blog, err := save.Save(r.Context())
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	if len(tags) > 0 {
		// Associate tags with the blog
		_, err = client.Blog.UpdateOneID(blog.ID).AddTags(tags...).Save(context.Background())
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, M{"error": "Failed to add tags"})
			return
		}
	}
	writeJSON(w, http.StatusOK, blog)
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	user_json := UserDetails{}

	if err := readJSON(w, r, &user_json); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": "Invalid Id received"})
		return
	}
	log.Printf("user_json: %v\n", user_json)

	// Fetch existing user to ensure it exists
	user, err := client.User.Get(r.Context(), id)
	fmt.Printf("user: %v\n", user)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// Handle the case where the entity was not found
			writeJSON(w, http.StatusNotFound, M{"error": "User not found"})
		} else {
			// Handle other errors
			writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		}
		return
	}

	// Apply partial updates
	update := client.User.UpdateOne(user)
	if user_json.Name != nil {
		update = update.SetName(*user_json.Name)
	}
	if user_json.Age != nil {
		update = update.SetAge(*user_json.Age)
	}
	log.Printf("user_json: %v\n", user_json)
	log.Printf("user_json.IsActive: %v\n", *user_json.IsActive)
	if user_json.IsActive != nil {
		update = update.SetIsActive(*user_json.IsActive)
	}

	updatedUser, err := update.Save(r.Context())
	log.Printf("updatedUser: %v\n", updatedUser)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, updatedUser)
}

func updateBlogById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	blog_json := BlogDetails{}

	if err := readJSON(w, r, &blog_json); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": "Invalid Id received"})
		return
	}
	log.Printf("blog_json: %v\n", blog_json)

	// Fetch existing user to ensure it exists
	blog, err := client.Blog.Get(r.Context(), id)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// Handle the case where the entity was not found
			writeJSON(w, http.StatusNotFound, M{"error": "Blog not found"})
		} else {
			// Handle other errors
			writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		}
		return
	}

	// Apply partial updates
	update := client.Blog.UpdateOne(blog)
	if blog_json.Title != nil {
		update = update.SetTitle(*blog_json.Title)
	}
	if blog_json.Description != nil {
		update = update.SetDescription(*blog_json.Description)
	}
	if blog_json.Episode != nil {
		update = update.SetEpisode(*blog_json.Episode)
	}
	if blog_json.UserId != nil {
		update = update.SetUserID(*blog_json.UserId)
	}
	log.Printf("blog_json: %v\n", blog_json)

	updated_blog, err := update.Save(r.Context())
	log.Printf("updated_blog: %v\n", updated_blog)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, updated_blog)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": "Invalid ID received"})
		return
	}

	err = client.User.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			writeJSON(w, http.StatusNotFound, M{"error": "User with ID " + id_string + " not found"})
		} else {
			writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		}
		return
	}

	writeJSON(w, http.StatusOK, M{"message": "User deleted successfully"})
}

func deleteByBlogId(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": "Invalid ID received"})
		return
	}

	err = client.Blog.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			writeJSON(w, http.StatusNotFound, M{"error": "Blog with ID " + id_string + " not found"})
		} else {
			writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		}
		return
	}

	writeJSON(w, http.StatusOK, M{"message": "User deleted successfully"})
}

func addFriendById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	var request FriendRequest
	log.Printf("request: %v\n", request)
	if err := readJSON(w, r, &request); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"message": err.Error()})
		return
	}
	log.Printf("request: %v\n", request)
	// Fetch users to validate existence
	user_entity := GetUserFromContext(r.Context())
	log.Printf("user_entity: %v\n", user_entity)

	friend, err := client.User.Get(r.Context(), request.FriendID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, M{"error": "Friend not found"})
		return
	}

	// Add friend to the user's friends list
	_, err = client.User.UpdateOne(user_entity).AddFriends(friend).Save(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, M{"message": "Friend added successfully"})
}

func deleteFriendById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	var request FriendRequest
	if err := readJSON(w, r, &request); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"message": err.Error()})
		return
	}

	// Fetch users to validate existence
	user := GetUserFromContext(r.Context())

	friend, err := client.User.Get(r.Context(), request.FriendID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, M{"error": "Friend not found"})
		return
	}

	// Remove friend from the user's friends list
	_, err = client.User.UpdateOne(user).RemoveFriends(friend).Save(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, M{"message": "Friend removed successfully"})
}

func getTags(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	var tags []struct {
		TagUpdateRequest
		ID    int `json:"id"`
		Count int `json:"blogs_count"`
	}
	err := client.Blog.Query().QueryTags().
		GroupBy(blog.FieldID, tag.Columns...).
		Aggregate(
			ent.As(ent.Count(), "blogs_count"),
		).
		Scan(r.Context(), &tags)
		// All(context.Background())
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, tags)
}

func updateTagById(w http.ResponseWriter, r *http.Request) {
	client := GetClient()

	tag_json := TagUpdateRequest{}
	if err := readJSON(w, r, &tag_json); err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	// Extract tag ID from URL path
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": "Invalid ID received"})
		return
	}

	// Fetch the existing tag
	tag_entity, err := client.Tag.Get(context.Background(), id)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// Handle the case where the entity was not found
			writeJSON(w, http.StatusNotFound, M{"error": "Tag not found"})
		} else {
			// Handle other errors
			writeJSON(w, http.StatusInternalServerError, M{"error": err.Error()})
		}
		return
	}

	// Apply partial updates
	update := client.Tag.UpdateOne(tag_entity)
	if tag_json.Name != nil {
		update = update.SetName(*tag_json.Name)
	}
	if tag_json.Type != nil {
		update = update.SetType(*tag_json.Type)
	}
	if tag_json.Category != nil {
		update = update.SetCategory(tag.Category(*tag_json.Category))
	}

	updatedTag, err := update.Save(context.Background())
	if err != nil {
		writeJSON(w, http.StatusBadRequest, M{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, updatedTag)
}
