package dto

type TodoParam struct {
	Name string `json:"name" binding:"required,max=255" example:"Eat Breakfast"`
}

type TodoUpdateParam struct {
	Name   string `json:"name" binding:"required,max=255" example:"Eat Breakfast"`
	IsDone bool   `json:"is_done" binding:"required" example:"true"`
}
