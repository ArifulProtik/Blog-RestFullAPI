package main

import (
	"fmt"
	"forum-api/db"
	"forum-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Dbprovider.PingDB()
	r := mux.NewRouter()
	routes.InitializeRoute(r)
	fmt.Println("Starting Server on 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
