package service

import "tudo-thrfting-clothes-server/internal/repo"

type ProductService struct {
	userRepo *repo.UserRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		userRepo: repo.NewUserRep(),
	}
}

func (us *ProductService) GetListUser() string {
	return us.userRepo.GetListUser()
}
