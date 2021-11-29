package routers

import (
	"myapp/middlewares"

	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.Use(middlewares.AuthMiddleware())

	// User
	router.POST("/user/register", controllers.UserController.Register)
	router.POST("/user/login", controllers.UserController.Login)
	router.GET("/user/me", controllers.UserController.Me)

	// Todo
	router.GET("/todos", controllers.TodoController.GetAll)
	router.GET("/todos/:id", controllers.TodoController.GetByID)
	router.POST("/todos", controllers.TodoController.Create)
	router.PUT("/todos/:id", controllers.TodoController.Update)
	router.DELETE("/todos/:id", controllers.TodoController.Delete)
}
