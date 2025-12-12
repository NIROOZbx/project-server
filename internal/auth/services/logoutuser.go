package services

import (
	"github.com/NIROOZbx/project-server/internal/auth/repositories"
)

func LogoutUser(userid uint) error {

	err:=repositories.UpdateTokenVersion(userid)

	if err!=nil{
		return err
	}

	return nil

}