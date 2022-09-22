package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kurianvarkey/go-todo-api/models"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

func (app App) GetTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetTodo")

	// Read dynamic id parameter; in the url we expect {id}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	todo := models.Todo{}

	// find the record from the todos table
	if result := app.DB.First(&todo, id); result.Error != nil { // also can be used result := app.DB.Where("id = ?", id).First(&todo)
		app.SendOutput(w, http.StatusBadRequest, "Todo not found with id: "+cast.ToString(id)+". Error: "+cast.ToString(result.Error))
		return
	}

	app.SendOutput(w, http.StatusOK, todo)
	return
}
