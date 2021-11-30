package dataloader

import (
	"myapp/models"
	services "myapp/services/db"
)

func TodoLoadByUserID(user *models.User) error {
	var (
		todos []*models.Todo
		err error
	)

	todos,err = services.Database.TodoGetByUserID(user.ID)
	if err != nil {
		return  err
	}

	user.Todos = todos

	return nil
}

func TodoLoadByUserIDs(users []*models.User) error {

	var (
		todos []*models.Todo
		err error
		userIDs []int
		todoMappedByUserID = map[int][]*models.Todo{}
	)

	for _,v := range users {
		userIDs = append(userIDs, v.ID)
	}

	todos,err = services.Database.TodoGetByWhereInUserIDs(userIDs)
	if err != nil {
		return  err
	}

	for _,v := range todos {
		todoMappedByUserID[v.UserID] = append(todoMappedByUserID[v.UserID], v)
	}

	for _,v := range users {
		v.Todos = todoMappedByUserID[v.ID]
	}

	return nil
}