package db

import (
	"github.com/tangren1998/ddd-cheanArch/user/domain"
)

type MysqlUserRepo struct {}

func (repo MysqlUserRepo) GetByID(id string) (*domain.User, error) {
	var user domain.User
	if result := db.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
