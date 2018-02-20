package repository

import "github.com/wmetaw/go-ddd-on-echo/domain"

// infrastructureで実装を期待する
type UserRepository interface {
	Get(id int) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(*domain.User) (*domain.User, error)
}
