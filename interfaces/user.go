package interfaces

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Me(*gin.Context)
}
