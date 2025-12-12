package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	userModel "github.com/NIROOZbx/project-server/internal/user/models"
	userRepo "github.com/NIROOZbx/project-server/internal/user/repositories"
	"gorm.io/gorm"
)

func GetAllUserService(page int,limit int)(*dtos.PaginationResponse,error){

	offset:=(page-1)*limit

	totalCount, err := userRepo.CountUsers()
    if err != nil {
        return nil, err
    }

	allUser,err:=userRepo.GetUsersWithPagination(limit,offset)
	if err!=nil{
		return nil,err
	}

	userDTO:=make([]userModel.UserProfile,0,len(allUser))

	for _,user:=range allUser{
		userDTO=append(userDTO, userModel.UserProfile{
			UserId: user.ID,
			Name: user.Name,
			Email: user.Email,
			IsVerified: user.IsVerified,
			ProfileImage: user.ProfileImage,
			IsBlocked: user.IsBlocked,
		})

	}
	totalPages := (totalCount + int64(limit) - 1) / int64(limit)


	paginationRes:=&dtos.PaginationResponse{
		Data: userDTO,
		TotalCount: totalCount,
		Page: page,
		TotalPages: totalPages,
	}

	return paginationRes,nil

}

func BlockUserService(userID string,adminId uint,blocked *bool) error{

	convertedUserId, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("eror in convertis",err)
		return errors.New("invalid user ID format") 
	}

	targetUserID := uint(convertedUserId)

	if targetUserID == adminId {
		return errors.New("cannot block your own account")
	}

	err=userRepo.BlockUserInDB(targetUserID,*blocked)
	if err != nil {
		if err==gorm.ErrRecordNotFound{
			return errors.New("user not found")
		}

		return err
	}
	return nil

}