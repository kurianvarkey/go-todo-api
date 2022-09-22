package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/kurianvarkey/go-todo-api/db"
	"github.com/kurianvarkey/go-todo-api/models"
)

func (app App) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllTodos")

	inputs, err := app.GetInputs(r)
	if err != nil {
		log.Println(err)
		return
	}

	var todos []models.Todo
	query := app.DB.Table("todos")

	if inputs.Get("from_date") != "" && inputs.Get("to_date") != "" {
		from_date, _ := time.Parse(date_layout, inputs.Get("from_date"))
		to_date, _ := time.Parse(date_layout, inputs.Get("to_date"))

		// I use some old tricks to substring the date only as the time parse will return date time
		query.Where("task_date BETWEEN ? AND ?", from_date.String()[0:10], to_date.String()[0:10])
	}

	if inputs.Get("search_title") != "" {
		query.Where("title LIKE ?", "%"+inputs.Get("search_title")+"%")
	}

	if inputs.Get("search_is_completed") == "1" {
		query.Where("is_completed = ?", inputs.Get("search_is_completed"))
	}

	//db_err := query.Find(&todos)
	pagination := db.Pagination{}
	db_err := query.Scopes(app.Paginate(r, query, todos, &pagination)).Order("id ASC").Find(&todos)
	if db_err.Error != nil {
		app.SendOutput(w, http.StatusBadRequest, "Error getting data")
		return
	}

	// setting the results as todos
	pagination.Results = todos

	// here we are sending the pagination struct as result
	app.SendOutput(w, http.StatusOK, pagination)
	return
}
