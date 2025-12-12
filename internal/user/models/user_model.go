package models

import "mime/multipart"

type UserProfile struct {
	UserId       uint   `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	IsVerified   bool   `json:"is_verified"`
	ProfileImage string `json:"profile_image"`
	IsBlocked    bool   `json:"is_blocked"`
}


type ResetPassword struct {
	OldPassword        string `json:"old_password" binding:"required,min=8"`
	NewPassword        string `json:"new_password" binding:"required,min=8"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required,min=8,eqfield=NewPassword"`
}

type ChangeName struct {
	Name string `json:"name" binding:"required"`
}

type UserProfileImage struct {
	ProfileImage *multipart.FileHeader `form:"image" binding:"required"`
}

type BlockUser struct {
	IsBlocked *bool `json:"is_blocked" binding:"required"`
}