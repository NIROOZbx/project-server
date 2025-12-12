package models

type SignInData struct {
	Name            string `form:"name" json:"name" binding:"required" `
	Password        string `form:"password" json:"password" binding:"required,min=8" `
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password" `
	Email           string `form:"email" json:"email" binding:"required,email" `
}

type LoginData struct {
	Email    string `form:"email" json:"email" binding:"required,email" `
	Password string `form:"password" json:"password" binding:"required,min=8" `
}
