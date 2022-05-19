package usecase

import "github.com/tangren1998/ddd-cheanArch/user/domain"

// UserUseCase 用户用例(实现 Use Case input Port，调用 Use Case Output Port)
type UserUseCase struct {
	Repo   domain.UserRepo
	Output UserOutput
}

// UserOutput - Use Case Output Port
type UserOutput interface {
	Standard(u *domain.User, err error)
}

func (u *UserUseCase) GetUserByID(id string) {
	user, err := u.Repo.GetByID(id)
	u.Output.Standard(user, err)
}
