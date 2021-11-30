package interfaces

import "myapp/models"

type DBInterface interface {

	// User
	UserCreate(models.User) (*models.User, error)
	UserGetByID(int) (*models.User, error)
	UserGetByEmail(string) (*models.User, error)
	UserGetAll() ([]*models.User, error)

	// Todo
	TodoCreate(models.Todo) (*models.Todo, error)
	TodoGetByID(int) (*models.Todo, error)
	TodoGetByUserID(int) ([]*models.Todo, error)
	TodoUpdate(models.Todo) (*models.Todo, error)
	TodoDelete(int) (*int, error)
	TodoGetByWhereInUserIDs([]int) ([]*models.Todo, error)
}
