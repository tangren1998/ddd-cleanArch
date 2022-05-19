package domain

// User - （实体）用户
type User struct {
	ID string
	Username string
}

func (User) TableName() string {
	return "user"
}

// UserRepo (存储库)
type UserRepo interface {
	GetByID(id string) (*User, error)
}
