package dtos

type ProductDTO struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Team        string  `json:"team"`
	League      string  `json:"league"`
	Season      int     `json:"season"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"totalCount"`
	Page       int         `json:"page"`
	TotalPages int64       `json:"totalPages"`
}

type PaginationInput struct {
	Page  int `form:"page,default=1"`
	Limit int `form:"limit,default=12"`
	Se string `form:"se"`
	SortBy    string `form:"sort_by,default=created_at"`
	Filter string `form:"filter"`
}
