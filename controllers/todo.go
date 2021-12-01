package controllers

import (
	"myapp/dto"
	"myapp/interfaces"
	"myapp/middlewares"
	"myapp/models"
	services "myapp/services/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	TodoController interfaces.TodoInterface
)

func init() {
	TodoController = new(todoController)
}

type todoController struct{}


// @Summary Get All todo attached to logged in user
// @Description Get All todo attached to logged in user, return list of todos
// @Tags todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 500
// @Router /api/todos [get]
func (c *todoController) GetAll(ctx *gin.Context) {

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	user, err := services.Database.UserGetByID(authorizedUser.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	todos, err := services.Database.TodoGetByUserID(user.ID)
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
		Data:   todos,
		Error:  "",
	})
}

func (c *todoController) GetByID(ctx *gin.Context) {

	todoId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	todo, err := services.Database.TodoGetByID(todoId)
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
		Data:   todo,
		Error:  "",
	})
}

func (c *todoController) Create(ctx *gin.Context) {

	var newTodo dto.TodoParam
	err := ctx.ShouldBind(&newTodo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	todo, err := services.Database.TodoCreate(models.Todo{
		Name:   newTodo.Name,
		IsDone: false,
		UserID: authorizedUser.ID,
	})

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
		Data:   todo,
		Error:  "",
	})
}

func (c *todoController) Update(ctx *gin.Context) {

	todoId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	var todo dto.TodoUpdateParam
	err = ctx.ShouldBind(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	tobeUpdatedTodo, err := services.Database.TodoGetByID(todoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if tobeUpdatedTodo.UserID != authorizedUser.ID {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "not allowed to updated other todo!",
		})
		return
	}

	newUpdatedTodo, err := services.Database.TodoUpdate(models.Todo{
		ID:     todoId,
		Name:   todo.Name,
		IsDone: todo.IsDone,
		UserID: authorizedUser.ID,
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
		Data:   newUpdatedTodo,
		Error:  "",
	})
}

func (c *todoController) Delete(ctx *gin.Context) {

	todoId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	tobeDeletedTodo, err := services.Database.TodoGetByID(todoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if tobeDeletedTodo.UserID != authorizedUser.ID {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "not allowed to updated other todo!",
		})
		return
	}

	_, err = services.Database.TodoDelete(todoId)
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
		Data:   todoId,
		Error:  "",
	})
}
