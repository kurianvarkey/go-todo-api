package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kurianvarkey/go-todo-api/models"

	"gorm.io/datatypes"
)

func (app App) AddTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("AddTodo")

	inputs, err := app.GetInputs(r)
	if err != nil {
		log.Println(err)
		return
	}

	/* if inputs.Get("title") == "" {
		fmt.Println("Title not found")
		return
	} */

	todo := models.Todo{} // same as var todo models.Todo

	todo.UserId, _ = strconv.Atoi(inputs.Get("user_id"))
	todo.Title = inputs.Get("title")
	todo.Description = inputs.Get("description")
	task_date, err := time.Parse(date_layout, inputs.Get("task_date"))
	if err != nil {
		log.Println(err)
	}
	todo.TaskDate = datatypes.Date(task_date)
	todo.IsCompleted, _ = strconv.ParseBool(inputs.Get("is_completed"))
	if inputs.Get("completed_date") != "" {
		completed_date, err := time.Parse(date_time_layout, inputs.Get("completed_date"))
		if err != nil {
			log.Println(err)
		}
		todo.CompletedDate = completed_date
	}

	// create a new records to the todo table
	if result := app.DB.Create(&todo); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	app.SendOutput(w, http.StatusCreated, todo)
	return
}
