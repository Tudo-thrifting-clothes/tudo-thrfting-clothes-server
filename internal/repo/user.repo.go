package repo

type UserRepo struct{}

func NewUserRep() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetListUser() string {
	return `Get list user`
}
