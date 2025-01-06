package model

type CreateProduct struct {
	Name        string `json:"name" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Stock       int64  `json:"stock" binding:"required"`
	Description string `json:"description" binding:"required"`
	BrandID     int64  `json:"brand_id"`
	CategoryID  int64  `json:"category_id"`
}
