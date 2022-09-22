// Routes will serve all the routes for this application
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kurianvarkey/go-todo-api/controllers"
	"github.com/kurianvarkey/go-todo-api/db"

	"github.com/gorilla/mux"
)

// This is the route function
func Routes(router *mux.Router) {
	app := controllers.NewDb(db.GetConnection())

	//This is the default route, we send back a welcome message
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to Kapi Todo, your IP is " + r.RemoteAddr)
	})

	router.HandleFunc("/todos", app.GetAllTodos).Methods(http.MethodGet) // Adding a Todo, HTTP Verb: Get
	// {id} is url parameter to send the id. It will be processed in the getTodo controller
	router.HandleFunc("/todos/{id}", app.GetTodo).Methods(http.MethodGet)       // Adding a Todo, HTTP Verb: Get
	router.HandleFunc("/todos", app.AddTodo).Methods(http.MethodPost)           // Adding a Todo, HTTP Verb: POST
	router.HandleFunc("/todos/{id}", app.UpdateTodo).Methods(http.MethodPut)    // Adding a Todo, HTTP Verb: Put
	router.HandleFunc("/todos/{id}", app.DeleteTodo).Methods(http.MethodDelete) // Adding a Todo, HTTP Verb: Delete
}
