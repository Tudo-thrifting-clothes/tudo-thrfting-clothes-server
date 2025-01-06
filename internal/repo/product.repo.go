package repo

import (
	"errors"
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

// GetListProduct retrieves a list of products from the database
func (pr *ProductRepo) GetListProduct() ([]po.Product, error) {
	var products []po.Product
	// Retrieve all products from the database
	if err := global.Mdb.Find(&products).Error; err != nil {
		// Return an empty slice and the error if it occurs
		return nil, err
	}
	// Return the list of products
	return products, nil
}

func (pr *ProductRepo) CreateProduct(input model.CreateProduct) (po.Product, error) {
	// Map model.CreateProduct to po.Product
	newProduct := po.Product{
		ProductName: input.Name,
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
