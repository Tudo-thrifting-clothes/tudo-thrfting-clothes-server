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

func (ps *ProductService) GetListProduct() ([]po.Product, error) {
	return ps.productRepo.GetListProduct()
}

func (ps *ProductService) CreateProduct(product model.CreateProduct) (po.Product, error) {
	return ps.productRepo.CreateProduct(product)
}
