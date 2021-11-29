package dto

type UserRegisterParam struct {
	Name     string `json:"name" binding:"required,min=5,max=255" example:"Louis Aldorio"`
	Email    string `json:"email" binding:"required,max=255,email" example:"louisaldorio@gmail.com"`
	Password string `json:"password"  binding:"required,min=4" example:"12345"`
}

type UserLoginParam struct {
	Email    string `json:"email" binding:"required,max=255,email" example:"louisaldorio@gmail.com"`
	Password string `json:"password"  binding:"required,min=4" example:"12345"`
}

type UserTokenResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}
