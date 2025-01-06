package repo

import (
	"errors"
	"fmt"
	"tudo-thrfting-clothes-server/global"
	"tudo-thrfting-clothes-server/internal/model"
	"tudo-thrfting-clothes-server/internal/po"
)

// ProductRepo struct
type ProductRepo struct{}

// NewProductRepo creates and returns a new ProductRepo instance
func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

// Get list of products
func (pr *ProductRepo) GetListProduct() ([]model.Products, error) {
	var productList []model.Products

	// Directly query the database into model.Products
	if err := global.Mdb.Model(&po.Product{}).
		Select("id, name, price, stock, slug, description").
		Find(&productList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	return productList, nil
}

// CreateProduct creates a new product
func (pr *ProductRepo) CreateProduct(input model.CreateProduct) (po.Product, error) {
	// Map model.CreateProduct to po.Product
	newProduct := po.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       int32(input.Price),
		Stock:       int32(input.Stock),
		// Map other fields as necessary
	}

	// Create a new product in the database
	if err := global.Mdb.Create(&newProduct).Error; err != nil {
		// Return an empty product and the error if it occurs
		return po.Product{}, errors.New("failed to create product: " + err.Error())
	}
	// Return the created product
	return newProduct, nil
}

// Create a new product brand
func (pr *ProductRepo) CreateProductBrand(input model.CreateProductBrand) (model.ProductBrand, error) {
	// Map model.CreateProductBrand to po.Brand
	newBrand := po.Brand{
		BrandName: input.BrandName,
		BrandNote: input.BrandNote,
	}

	// Create a new brand in the database
	if err := global.Mdb.Create(&newBrand).Error; err != nil {
		return model.ProductBrand{}, errors.New("failed to create brand: " + err.Error())
	}

	// Map po.Brand to model.ProductBrand
	createdBrand := model.ProductBrand{
		ID:        int(newBrand.ID),
		BrandName: newBrand.BrandName,
		BrandNote: newBrand.BrandNote,
	}

	// Return the created brand
	return createdBrand, nil
}
