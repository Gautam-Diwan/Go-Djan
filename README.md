
# Go Djan

Go Djan is a simple web service project that provides CRUD APIs for managing blogs, users, tags, and friendships. The project is my attempt as a Django developer and is developed using the Go programming language.

## Overview

This project is designed to help developers familiar with Django to explore how similar functionalities can be implemented in Go. It covers basic user authentication, blog management, and tag management. The project uses PASETO for token-based authentication and includes handlers for managing users, blogs, and tags.

## Features

- **User Management**: 
  - User signup and login with bcrypt for password hashing.
  - Update user information and manage friends.

- **Blog Management**: 
  - Create, read, update, and delete blog posts.
  - Associate tags with blog posts.
  
- **Tag Management**:
  - Create, update, and retrieve tags associated with blog posts.
  
- **Friendship Management**: 
  - Add and remove friends.

## Files

### `auth.go`

- Handles authentication logic using PASETO tokens.
- Provides functions to retrieve the user from the context.

### `handlers.go`

- Contains handlers for all CRUD operations related to users, blogs, tags, and friends.
- Includes the following functionalities:
  - `getUsers()`: Fetch all users with their blogs and friends.
  - `getBlogs()`: Fetch all blogs, with optional filtering by tag category.
  - `getUserById()`: Fetch a user by their ID.
  - `getBlogById()`: Fetch a blog by its ID.
  - `loginHandler()`: Handle user login and return a PASETO token.
  - `signUpHandler()`: Handle user signup.
  - `signOutHandler()`: Handle user logout by invalidating the token.
  - `createBlog()`: Create a new blog post.
  - `updateUserById()`: Update user information.
  - `updateBlogById()`: Update a blog post.
  - `deleteUserById()`: Delete a user by ID.
  - `deleteByBlogId()`: Delete a blog post by ID.
  - `addFriendById()`: Add a friend by user ID.
  - `deleteFriendById()`: Remove a friend by user ID.
  - `getTags()`: Retrieve all tags with their associated blog count.
  - `updateTagById()`: Update tag details.

### `helpers.go`

- Utility functions for handling JSON encoding/decoding and HTTP responses.
- Provides functions to:
  - Write JSON responses.
  - Handle malformed requests.
  - Load environment variables using viper.

## Getting Started

### Prerequisites

- Go 1.22 or higher.
- A PostgreSQL database (or any other supported by Ent).

### Setup

1. Clone the repository:

    ```bash
    git clone git@github.com:Gautam-Diwan/Go-Djan.git
    cd go-djan
    ```

2. Set up the environment variables by creating a `.env` file:

    ```
    DATABASE_URL=your-database-url
    PASETO_KEY=your-paseto-key
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

    Alternatively, you can use air for live reloading
    ```bash
    air init
    air bench
    ```

4. The API will be available at `http://localhost:8080`.

### API Endpoints

- `GET /users`: Retrieve all users.
- `GET /users/{id}`: Retrieve a user by ID.
- `POST /users/login`: Log in a user.
- `POST /users/signup`: Sign up a new user.
- `PUT /users/{id}`: Update a user's information.
- `DELETE /users/{id}`: Delete a user.
- `GET /blogs`: Retrieve all blogs.
- `GET /blogs/{id}`: Retrieve a blog by ID.
- `POST /blogs`: Create a new blog post.
- `PUT /blogs/{id}`: Update a blog post.
- `DELETE /blogs/{id}`: Delete a blog post.
- `GET /tags`: Retrieve all tags.
- `PUT /tags/{id}`: Update a tag's information.
- `POST /friends`: Add a friend.
- `DELETE /friends`: Remove a friend.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
