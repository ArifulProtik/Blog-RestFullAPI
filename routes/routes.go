package routes

import (
	"forum-api/controller"

	"github.com/gorilla/mux"
)

// InitializeRoute Creates All the Route
func InitializeRoute(r *mux.Router) {
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/", controller.Home).Methods("OPTION")
	r.HandleFunc(("/api/v1/signup", controller.Signup).Methods("POST"))
	r.HandleFunc(("/api/v1/signup", sicontroller.Signup).Methods("OPTION"))
}
