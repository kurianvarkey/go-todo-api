package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kurianvarkey/go-todo-api/models"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"gorm.io/datatypes"
)

func (app App) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateTodo")

	inputs, err := app.GetInputs(r)
	if err != nil {
		log.Println(err)
		return
	}

	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	todo := models.Todo{}

	// find the record from the todos table
	if result := app.DB.First(&todo, id); result.Error != nil { // also can be used result := app.DB.Where("id = ?", id).First(&todo)
		app.SendOutput(w, http.StatusBadRequest, "Todo not found with id: "+cast.ToString(id)+". Error: "+cast.ToString(result.Error))
		return
	}

	todo.UserId, _ = strconv.Atoi(inputs.Get("user_id"))
	todo.Title = inputs.Get("title")
	todo.Description = inputs.Get("description")
	task_date, err := time.Parse(date_layout, inputs.Get("task_date"))
	if err != nil {
		log.Println(err)
	}
	todo.TaskDate = datatypes.Date(task_date)
	todo.IsCompleted, _ = strconv.ParseBool(inputs.Get("is_completed"))
	completed_date, err := time.Parse(date_time_layout, inputs.Get("completed_date"))
	if err != nil {
		log.Println(err)
	}
	todo.CompletedDate = completed_date

	// saves the todo row
	app.DB.Save(&todo)

	app.SendOutput(w, http.StatusOK, todo)
	return
}
