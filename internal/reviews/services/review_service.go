package services

import (
	"errors"
	"strconv"

	reviewModel "github.com/NIROOZbx/project-server/internal/reviews/models"
	reviewRepo "github.com/NIROOZbx/project-server/internal/reviews/respositories"
)

func AddProductReviewService(userID uint, productIdStr string,reviewData string) error {
	productID, err := strconv.Atoi(productIdStr)

	if err != nil {
		return  errors.New("invalid product ID format")
	}
	if !reviewRepo.HasUserPurchasedProduct(userID, uint(productID)) {
        return errors.New("you can only review products you have purchased")
    }

	newReview := &reviewModel.Review{
		UserID: userID,
		ProductID: uint(productID),
		Comment: reviewData,
	}

	err=reviewRepo.AddProductReviewInDB(newReview)

	if err != nil {
        return err
    }

	return nil

}


func GetProductReviewsService(productIDStr string) ([]reviewModel.ReviewResponse, error) {


	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	// 2. Fetch Data
	reviews, err := reviewRepo.GetReviewsByProductID(uint(productID))
	if err != nil {
		return nil, err
	}
	var response []reviewModel.ReviewResponse

    for _, r := range reviews {
        dto := reviewModel.ReviewResponse{
            ID:        r.ID,
            Comment:   r.Comment,
            CreatedAt: r.CreatedAt,
            UserName:  r.User.Name,
			UserImage: r.User.ProfileImage,
        }
        response = append(response, dto)
    }

    return response, nil

}


