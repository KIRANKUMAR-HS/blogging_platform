package apihandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/KIRANKUMAR-HS/blogging_platform/internal/authservice"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/model"
	db "github.com/KIRANKUMAR-HS/blogging_platform/internal/psql"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/utils"
	"github.com/gorilla/mux"
)

type Bloghandler struct {
	db *db.PsqlClient
	a  *authservice.AuthService
}

func NewBlogServer(db *db.PsqlClient, auth *authservice.AuthService) (*Bloghandler, error) {
	return &Bloghandler{
		db: db,
		a:  auth,
	}, nil
}

func (b *Bloghandler) Start(Adress string, r *mux.Router) error {
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
		return err
	}
	return nil
}

// GetAllPosts retrieves all posts, with optional filtering and pagination
func (bg *Bloghandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	offsetParam := r.URL.Query().Get("offset")

	// Default values for pagination
	limit := 10
	offset := 0

	// Parse pagination parameters
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}

	if offsetParam != "" {
		if o, err := strconv.Atoi(offsetParam); err == nil {
			offset = o
		}
	}
	posts, err := bg.db.GetPosts(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// GetPost retrieves a single post by its ID
func (bg *Bloghandler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := bg.db.GetPostByID(int64(postID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// CreatePost creates a new blog post
func (bg Bloghandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	if _, err := bg.db.CreatePost(&newPost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

// UpdatePost updates an existing post
func (bg *Bloghandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updatedPost model.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	updatedPost.ID = int64(postID)

	if err := bg.db.UpdatePost(&updatedPost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePost deletes a post by its ID
func (bg *Bloghandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	//extracting id for filtering
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := bg.db.DeletePost(int64(postID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// GetAllPostsByfiltering gives posts based on auther or filtering
func (bg *Bloghandler) GetAllPostsByFiltering(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters for filtering
	author := r.URL.Query().Get("author")
	createdAfter := time.Time{}

	// Convert the `created_after` parameter to a `time.Time`
	createdAfterStr := r.URL.Query().Get("created_after")
	if createdAfterStr != "" {
		var err error
		createdAfter, err = time.Parse(time.RFC3339, createdAfterStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
	}

	// Set default values for pagination
	limit := 10
	offset := 0

	// Parse pagination parameters
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	// Retrieve posts with filtering and pagination
	posts, err := bg.db.GetAllPostsByfiltering(author, createdAfter, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the list of posts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// user login
func (h *Bloghandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Printf(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.a.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// CreateUser will creates a new user register
func (bg Bloghandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Check if the role is populated correctly

	password_hash, err := utils.HashPassword(user.Password_hash)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusBadRequest)
		return
	}

	user.Password_hash = password_hash

	if _, err := bg.db.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
