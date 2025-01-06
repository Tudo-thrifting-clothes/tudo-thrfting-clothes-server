package service

import (
	"tudo-thrfting-clothes-server/internal/model"
	"tudo-thrfting-clothes-server/internal/po"
	"tudo-thrfting-clothes-server/internal/repo"
)

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repo.NewProductRepo(),
	}
}

func (ps *ProductService) GetListProduct() ([]model.Products, error) {
	return ps.productRepo.GetListProduct()
}

func (ps *ProductService) CreateProduct(product model.CreateProduct) (po.Product, error) {
	// check product name exist
	return ps.productRepo.CreateProduct(product)
}

func (ps *ProductService) CreateProductBrand(brand model.CreateProductBrand) (model.ProductBrand, error) {
	// check product name exist
	return ps.productRepo.CreateProductBrand(brand)
}
