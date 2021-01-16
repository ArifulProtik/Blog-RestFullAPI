package routes

import (
	"forum-api/controller"

	"github.com/gorilla/mux"
)

// InitializeRoute Creates All the Route
func InitializeRoute(r *mux.Router) {
	r.HandleFunc("/", controller.Home).Methods("GET", "OPTION")

	r.HandleFunc("/signup", controller.JSONandCORS(controller.Signup)).Methods("POST", "OPTION")
	r.HandleFunc("/signin", controller.JSONandCORS(controller.Signin)).Methods("POST", "OPTION")

	r.HandleFunc("/posts", controller.JSONandCORS(controller.GetAllPost)).Methods("GET", "OPTION")
	r.HandleFunc("/post/{slug}", controller.JSONandCORS(controller.Singlepost)).Methods("GET", "OPTION")
	r.HandleFunc("/comments/{slug}", controller.JSONandCORS(controller.GetComments)).Methods("GET", "OPTION")

	// Authorized Route
	r.HandleFunc("/create", controller.JSONandCORS(controller.IsAuth(controller.CreatePost))).Methods("POST", "OPTION")
	r.HandleFunc("/post", controller.JSONandCORS(controller.IsAuth(controller.UpdatePost))).Methods("PUT", "OPTION")
	r.HandleFunc("/post/{id}", controller.JSONandCORS(controller.IsAuth(controller.DeletePost))).Methods("DELETE", "OPTION")
	r.HandleFunc("/createcomment", controller.JSONandCORS(controller.IsAuth(controller.SaveComment))).Methods("POST", "OPTION")
	r.HandleFunc("/comment/{uid}", controller.JSONandCORS(controller.IsAuth(controller.DeleteComment))).Methods("DELETE", "OPTION")

}
