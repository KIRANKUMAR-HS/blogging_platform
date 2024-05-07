package router

import (
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/apihandler"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/authservice"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/middleware"
	"github.com/gorilla/mux"
	// "net/http"
)

// NewRouter creates a new Gorilla Mux router and defines routes.
func NewRouter(h *apihandler.Bloghandler, a *authservice.AuthService) (*mux.Router, error) {

	router := mux.NewRouter() // Create a new router
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", h.Login).Methods("POST")
	auth.HandleFunc("/register", h.CreateUser).Methods("POST")

	// Define the routes and map them to handlers
	PostRouter := router.PathPrefix("/posts").Subrouter()
	PostRouter.Use(middleware.AuthMiddleware(a.SecretKey))
	router.Use()
	PostRouter.HandleFunc("", h.GetAllPosts).Methods("GET")
	PostRouter.HandleFunc("/{id}", h.GetPost).Methods("GET")
	PostRouter.HandleFunc("/paging/", h.GetAllPostsByFiltering).Methods("GET")

	AdminRouter := router.PathPrefix("/admin/posts").Subrouter()
	AdminRouter.Use(middleware.AdminOnlyMiddleware(a.SecretKey))
	AdminRouter.HandleFunc("", h.CreatePost).Methods("POST")
	AdminRouter.HandleFunc("/{id}", h.UpdatePost).Methods("PUT")
	AdminRouter.HandleFunc("/{id}", h.DeletePost).Methods("DELETE")
	// for role based Authentication
	// postRoutes.HandleFunc("", h.CreatePost).Methods("POST").Use(middleware.AdminOnlyMiddleware)
	// postRoutes.HandleFunc("/{id}", h.UpdatePost).Methods("PUT").Use(middleware.AdminOnlyMiddleware)
	// PostRouter.HandleFunc("/{id}", h.DeletePost).Methods("DELETE").Subrouter().Use(middleware.AdminOnlyMiddleware)

	return router, nil // Return the configured router
}
