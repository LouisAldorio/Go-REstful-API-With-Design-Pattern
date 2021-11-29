package interfaces

import "github.com/gin-gonic/gin"

type TodoInterface interface {
	Create(*gin.Context)
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
