package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     int
	Title  string
	Status bool
}

var i = 0

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {

	x := c.PostForm("todo")
	var todo Todo
	todo.ID = i
	i++
	todo.Title = x
	todo.Status = true

	c.JSON(http.StatusOK, todo)

}
