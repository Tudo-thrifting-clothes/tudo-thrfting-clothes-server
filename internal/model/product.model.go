package model

type CreateProduct struct {
	Name              string `json:"name" binding:"required" example:"product name"`
	Price             int64  `json:"price" binding:"required" example:"100000"`
	Stock             int64  `json:"stock" binding:"required" example:"100"`
	Description       string `json:"description" binding:"required" example:"product description"`
	ProductBrandID    int64  `json:"brand_id" example:"1"`
	ProductCategoryID int64  `json:"category_id" example:"1"`
}

type Products struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type ProductBrand struct {
	ID        int    `json:"id"`
	BrandName string `json:"brand_name"`
	BrandNote string `json:"brand_note"`
}

type ProductCategory struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
	CategoryNote string `json:"category_note"`
}

type CreateProductBrand struct {
	BrandName string `json:"brand_name" binding:"required" example:"brand name"`
	BrandNote string `json:"brand_note" example:"brand note"`
}

type CreateProductCategory struct {
	CategoryName string `json:"category_name" binding:"required" example:"category name"`
	CategoryNote string `json:"category_note" example:"category note"`
}
