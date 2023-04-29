package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"practice-todo-go-api/app"

	"github.com/revel/revel"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoController struct {
	*revel.Controller
}

/** タスク作成 */
func (c TodoController) Create() revel.Result {
	var todo Todo

	err := json.NewDecoder(c.Request.GetBody()).Decode(&todo)
	if err != nil {
		return c.RenderJSON(revel.ErrorResult{Error: errors.New(strconv.Itoa(http.StatusBadRequest) + "Invalid request body")})
	}

	query := "INSERT INTO todos (title, completed) VALUES (?, ?)"
	result, err := app.DB.Exec(query, todo.Title, todo.Completed)
	if err != nil {
		return c.RenderJSON(revel.ErrorResult{Error: errors.New(strconv.Itoa(http.StatusInternalServerError) + "Error inserting todo into the database")})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.RenderJSON(revel.ErrorResult{Error: errors.New((strconv.Itoa(http.StatusInternalServerError) + "Error getting the last inserted ID"))})
	}

	todo.ID = int(id)
	return c.RenderJSON(todo)
}

