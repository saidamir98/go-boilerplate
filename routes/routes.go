package routes

import (
	"github.com/gorilla/mux"
	"github.com/saidamir98/go-boilerplate/controllers"
	"github.com/saidamir98/go-boilerplate/middlewares"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/api", controllers.Index).Methods("GET")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middlewares.JwtVerify)

	// s.HandleFunc("/my-posts", controllers.ListUserPosts).Methods("GET")
	// s.HandleFunc("/posts", controllers.CreatPost).Methods("POST")
	// s.HandleFunc("/posts", controllers.ListPosts).Methods("GET")
	// s.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	// s.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	// s.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	// s.HandleFunc("/comments", controllers.CreatComment).Methods("POST")
	// s.HandleFunc("/comments/{id}", controllers.UpdateComment).Methods("PUT")
	// s.HandleFunc("/comments/{id}", controllers.DeleteComment).Methods("DELETE")

	// s.HandleFunc("/replies", controllers.CreatReply).Methods("POST")
	// s.HandleFunc("/replies/{id}", controllers.UpdateReply).Methods("PUT")
	// s.HandleFunc("/replies/{id}", controllers.DeleteReply).Methods("DELETE")
	return r
}
