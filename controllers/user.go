package controllers

import (
	"myapp/interfaces"
	"myapp/middlewares"
	"myapp/models"
	"net/http"
	"time"

	"myapp/dto"

	services "myapp/services/db"

	"github.com/gin-gonic/gin"
	"myapp/dataloader"
)

var (
	UserController interfaces.UserController
)

func init() {
	UserController = new(userController)
}

type userController struct{}


// @Summary Register New User
// @Description Register new user to the system
// @Tags user
// @Accept json
// @Produce json
// @Param AdminData body dto.UserRegisterParam true "User registration Data"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 500
// @Router /api/user/register [post]
func (c *userController) Register(ctx *gin.Context) {

	var input dto.UserRegisterParam
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	tx := services.BeginTransaction()

	_, err = services.Database.UserGetByEmail(input.Email)
	if err == nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Email already exists!",
		})
		return
	}

	user, err := services.Database.UserCreate(models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data: dto.UserTokenResponse{
			Token: middlewares.JwtGenerate(user.ID),
		},
		Error: "",
	})
}

func (c *userController) Login(ctx *gin.Context) {

	var input dto.UserLoginParam
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	tx := services.BeginTransaction()

	user, err := services.Database.UserGetByEmail(input.Email)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if user.Password != input.Password {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Invalid Password!",
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data: dto.UserTokenResponse{
			Token: middlewares.JwtGenerate(user.ID),
		},
		Error: "",
	})
}

func (c *userController) Me(ctx *gin.Context) {

	tx := services.BeginTransaction()

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	user, err := services.Database.UserGetByID(authorizedUser.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	 err = dataloader.TodoLoadByUserID(user)
	 if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   user,
		Error:  "",
	})
}

func (c *userController) GetAll(ctx *gin.Context) {

	tx := services.BeginTransaction()

	users, err := services.Database.UserGetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	err = dataloader.TodoLoadByUserIDs(users)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   users,
		Error:  "",
	})
}