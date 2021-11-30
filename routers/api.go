package routers

import (
	"myapp/middlewares"

	"myapp/controllers"


	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.Use(middlewares.AuthMiddleware())

	apiRouter := router.Group("/api")

	// User
	apiRouter.POST("/user/register", controllers.UserController.Register)
	apiRouter.POST("/user/login", controllers.UserController.Login)
	apiRouter.GET("/user/me", controllers.UserController.Me)

	apiRouter.GET("/users", controllers.UserController.GetAll)

	// Todo
	apiRouter.GET("/todos", controllers.TodoController.GetAll)
	apiRouter.GET("/todos/:id", controllers.TodoController.GetByID)
	apiRouter.POST("/todos", controllers.TodoController.Create)
	apiRouter.PUT("/todos/:id", controllers.TodoController.Update)
	apiRouter.DELETE("/todos/:id", controllers.TodoController.Delete)
}
