// package tests

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	// "blogging-platform/controllers"
// 	// "blogging-platform/models"

// 	"github.com/gorilla/mux"
// 	"github.com/KIRANKUMAR-HS/blogging_platform/internal/apihandler"
// 	// db "github.com/KIRANKUMAR-HS/blogging_platform/internal/psql"
// )

// func TestCreatePost(t *testing.T) {
// 	// Db := db.PsqlClient{}
// 	postController := &apihandler.Bloghandler{}
// 	router := mux.NewRouter()

// 	router.HandleFunc("/posts", postController.CreatePost).Methods("POST")

// 	body := `{"title": "Test Post", "content": "This is a test post.", "author": "Tester"}`
// 	req, err := http.NewRequest("POST", "/posts", strings.NewReader(body))
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v", err)
// 	}

// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)

// 	if rec.Code != http.StatusCreated {
// 		t.Fatalf("Expected status code %d but got %d", http.StatusCreated, rec.Code)
// 	}

// 	expected := `{"title": "Test Post", "content": "This is a test post.", "author": "Tester"}`
// 	if rec.Body.String() != expected {
// 		t.Fatalf("Expected response body to be %s but got %s", expected, rec.Body.String())
// 	}

// 	// Set up the router
// 	// router.HandleFunc("/posts", postController.CreatePost).Methods("POST")
// 	// router.HandleFunc("/posts", postController.GetAllPosts).Methods("GET")
// 	// router.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
// 	// router.HandleFunc("/posts/{id}", postController.UpdatePost).Methods("PUT")
// 	// router.HandleFunc("/posts/{id}", postController.DeletePost).Methods("DELETE")

// 	// return router, db
// }

package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "blogging-platform/controllers"
	// "blogging-platform/models"

	"github.com/KIRANKUMAR-HS/blogging_platform/internal/apihandler"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/model"
	"github.com/gorilla/mux"
)

// Helper function to create a test server
func setupTestServer() *mux.Router {
	// Mock database connection (in-memory or test-specific setup)
	// db := &apihandler.Bloghandler{}
	// Use the same database connection logic as in your production code
	// db.Connect("your_test_db_connection_string")

	// Create controller instances
	postController := &apihandler.Bloghandler{}

	// Set up the router
	router := mux.NewRouter()
	router.HandleFunc("/posts", postController.CreatePost).Methods("POST")
	router.HandleFunc("/posts", postController.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", postController.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", postController.DeletePost).Methods("DELETE")

	return router
}

func TestCreatePost(t *testing.T) {
	router := setupTestServer()

	body := `{"title": "Test Post", "content": "This is a test post.", "author": "Tester"}`
	req, err := http.NewRequest("POST", "/posts", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var createdPost model.Post
	if err := json.Unmarshal(rec.Body.Bytes(), &createdPost); err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	if createdPost.Title != "Test Post" || createdPost.Content != "This is a test post." {
		t.Fatalf("Unexpected post data: %v", createdPost)
	}
}

func TestGetAllPosts(t *testing.T) {
	router := setupTestServer()

	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	// Optionally, you can check the response data to ensure it contains the expected posts
}

func TestGetPost(t *testing.T) {
	router := setupTestServer()

	// Assume we have a post with ID 1
	req, err := http.NewRequest("GET", "/posts/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	// Further checks can be done on the returned post
}

func TestUpdatePost(t *testing.T) {
	router := setupTestServer()

	// Assume we're updating a post with ID 1
	body := `{"title": "Updated Title", "content": "Updated content"}`
	req, err := http.NewRequest("PUT", "/posts/1", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	// You can check if the updated post contains the correct information
}

func TestDeletePost(t *testing.T) {
	router := setupTestServer()

	// Assume we're deleting a post with ID 1
	req, err := http.NewRequest("DELETE", "/posts/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("Expected status code %d but got %d", http.StatusNoContent, rec.Code)
	}

	// Further checks can be made to ensure the post was actually deleted
}
