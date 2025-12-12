package models

type ForgotPassword struct {
    NewPassword        string `json:"new_password" binding:"required,min=8"`
    ConfirmNewPassword string `json:"confirm_new_password" binding:"required,min=8,eqfield=NewPassword"`
}

