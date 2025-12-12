package services

import (
	"errors"
	

	"strconv"

	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	productRepo "github.com/NIROOZbx/project-server/internal/products/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"
)



func ShowSingleProductService(productId string) (*dtos.ProductDTO,error){

	convertedProductId, err := strconv.Atoi(productId)
    
   
	if err != nil {
		return nil, errors.New("invalid product ID format") 
	}
	
	product,err:=productRepo.GetSingleProductFromDB(uint(convertedProductId))


	if err != nil {
		return nil, err
	}

	if product.Stock <= 0 {
        return nil, errors.New("product is currently out of stock")
    }

	productDTO:=&dtos.ProductDTO{
			Id:          product.ID,
			Name:        product.Name,
			Team:        product.Team,
			League:      product.League,
			Season:      product.Season,
			Stock:       product.Stock,
			Price:       product.Price,
			Currency:    product.Currency,
			Image:       product.Image,
			Category:    product.Category,
			Description: product.Description,
	}

	return productDTO, nil
}

func GetPaginationResults(input *dtos.PaginationInput) (*dtos.PaginationResponse,error){

	offset:=(input.Page-1)*input.Limit

	totalCount, err := productRepo.CountAllProducts()
	if err != nil {
		return nil, err
	}

	if totalCount == 0 {
		return nil, nil
	}

	products,err:=productRepo.GetPaginationFromDB(offset,input)
	if err != nil {
		return nil, err
	}


	allProducts:=mapToDTO(products)

	totalPages := (totalCount + int64(input.Limit) - 1) / int64(input.Limit)

	


	paginationResponse :=&dtos.PaginationResponse{
		Data: allProducts,
		TotalPages: totalPages,
		TotalCount: totalCount,
		Page: input.Page,

	}

	return paginationResponse,nil


}

func mapToDTO(products []productModel.Product) []dtos.ProductDTO {

	productsDTO := make([]dtos.ProductDTO, 0, len(products))

	for _, p := range products {
		productsDTO = append(productsDTO, dtos.ProductDTO{
			Id:          p.ID,
			Name:        p.Name,
			Team:        p.Team,
			League:      p.League,
			Season:      p.Season,
			Stock:       p.Stock,
			Price:       p.Price,
			Currency:    p.Currency,
			Image:       p.Image,
			Category:    p.Category,
			Description: p.Description,
			
		})
	}

	return productsDTO
}

func GetHomePageProducts() ([]dtos.ProductDTO, error) {

    products,err:= productRepo.GetLimitedProducts(12)
	allProducts:=mapToDTO(products)

	return allProducts,err
}