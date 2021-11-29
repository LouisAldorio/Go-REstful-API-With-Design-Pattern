package controllers

import (
	"fmt"
	"myapp/interfaces"
	"myapp/middlewares"
	"myapp/models"
	"net/http"
	"time"

	"myapp/dto"

	services "myapp/services/db"

	"github.com/gin-gonic/gin"
)

var (
	UserController interfaces.UserController
)

func init() {
	UserController = new(userController)
}

type userController struct{}

func (c *userController) Register(ctx *gin.Context) {

	tx := services.BeginTransaction()

	var input dto.UserRegisterParam
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	user, err := services.Database.UserCreate(models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.UserTokenResponse{
		Token: middlewares.JwtGenerate(user.ID),
	})
}

func (c *userController) Login(ctx *gin.Context) {

	tx := services.BeginTransaction()

	var input dto.UserLoginParam
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	user, err := services.Database.UserGetByEmail(input.Email)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if user.Password != input.Password {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.UserTokenResponse{
		Token: middlewares.JwtGenerate(user.ID),
	})
}

func (c *userController) Me(ctx *gin.Context) {

	tx := services.BeginTransaction()

	user := middlewares.AuthCtx(ctx.Request.Context())
	fmt.Println(user)

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
