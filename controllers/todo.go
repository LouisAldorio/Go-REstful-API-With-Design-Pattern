package controllers

import (
	"myapp/interfaces"

	"github.com/gin-gonic/gin"
)

var (
	TodoController interfaces.TodoInterface
)

func init() {
	TodoController = new(todoController)
}

type todoController struct{}

func (controller *todoController) GetAll(c *gin.Context) {
	
}

func (controller *todoController) GetByID(c *gin.Context) {
	
}

func (controller *todoController) Create(c *gin.Context) {
	
}

func (controller *todoController) Update(c *gin.Context) {
	
}

func (controller *todoController) Delete(c *gin.Context) {
	
}
