package services

import (
	"context"
	"errors"
	"strconv"

	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	productRepo "github.com/NIROOZbx/project-server/internal/products/repositories"
	
	fileupload "github.com/NIROOZbx/project-server/internal/shared/fileUpload"
	"gorm.io/gorm"
)

func AddProductService(ctx context.Context, newProduct *productModel.AddProductInput) error {

	imageURL, err := fileupload.UploadFileToCloudinary(ctx, newProduct.ImageFile)
	if err != nil {
		return errors.New("image upload failed")
	}

	newProductToAddInDB := &productModel.Product{
		Name:        newProduct.Name,
		Team:        newProduct.Team,
		League:      newProduct.League,
		Season:      newProduct.Season,
		Stock:       newProduct.Stock,
		Price:       newProduct.Price,
		Currency:    newProduct.Currency,
		Image:       imageURL,
		Category:    newProduct.Category,
		Description: newProduct.Description,
	}

	return productRepo.AddProductInDB(newProductToAddInDB)
}

func DeleteProductService(productId string) error {

	convertedProdId, err := strconv.Atoi(productId)

	if err != nil {
		return errors.New("invalid product ID format")
	}

	err = productRepo.RemoveProductFromDB(uint(convertedProdId))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("couldn't find the product")
		}

		return err
	}

	return nil

}


func UpdateProductService(ctx context.Context,productIdStr string,input *productModel.UpdateProductInput) error{

	productID, err := strconv.Atoi(productIdStr)
	if err != nil {
		return errors.New("invalid query format")
	}
	product, err := productRepo.GetSingleProductFromDB(uint(productID))
	if err != nil {
		return err 
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	
	if input.Team != "" {
		product.Team = input.Team
	}

	if input.League != "" {
		product.League = input.League
	}

	if input.Season != 0 {
		product.Season = input.Season
	}

	if input.Price > 0 {
		product.Price = input.Price
	}

	if input.Stock != nil {
    if *input.Stock >= 0 {
        product.Stock = *input.Stock
    }
}

	if input.Category != "" {
		product.Category = input.Category
	}

	if input.Description != "" {
		product.Description = input.Description
	}

	if input.Currency != "" {
		product.Currency = input.Currency
	}


	if input.ImageFile != nil {
		newImageURL, err := fileupload.UploadFileToCloudinary(ctx,input.ImageFile)
		if err != nil {
			return errors.New("failed to upload new image")
		}

		product.Image = newImageURL
	}


	return productRepo.UpdateProductInDB(product)


}
