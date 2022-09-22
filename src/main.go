// This is the main package, we start from here
//
// My aim is to create a simple todo list. This example is done without any frameworks and used standard packages
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kurianvarkey/go-todo-api/db"
	"github.com/kurianvarkey/go-todo-api/routes"
)

// init will trigger the start of the application before the main
func init() {
	//we will just print on cmd about started
	log.Println("Kapi Todos API v1.0.0 is started")

	//we will load our env file settings
	error_env := godotenv.Load()
	if error_env != nil {
		panic("Failed to load the env file, please check")
	}

	log.Println("env settings loaded")
}

// main function will trigger at the start of application after the init call
func main() {
	defer db.Disconnect()

	log.Println("Kapi Todos running routes")

	//we will call the routes to trigger
	router := mux.NewRouter()
	routes.Routes(router)

	// start the http server listening to the port
	http.ListenAndServe(os.Getenv("PORT"), router)
}
