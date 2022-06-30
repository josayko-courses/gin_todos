package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.JSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	todo.Completed = !todo.Completed
	context.JSON(http.StatusOK, todo)

}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index, err := getTodoIndexById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	todos = append(todos[:index], todos[index+1:]...)
	context.JSON(http.StatusOK, gin.H{"message": err.Error()})
}

func getTodoIndexById(id string) (int, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("todo not found")
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
