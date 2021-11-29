package dto

type TodoParam struct {
	Name string `json:"name" binding:"required,max=255" example:"Eat Breakfast"`
}