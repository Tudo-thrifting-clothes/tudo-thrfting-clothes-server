package repo

type ProductRepo struct{}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (ur *ProductRepo) GetListUser() string {
	return `Get list user`
}
