package application

import (
	"github.com/wmetaw/go-ddd-on-echo/domain"
	"github.com/wmetaw/go-ddd-on-echo/infrastructure/persistence"
)

// GetUsers returns user list
func Users() ([]*domain.User, error) {

	return persistence.NewUserRepositoryWithRDB().GetAll()
}

// GetUser returns user list
func UsersGet(id int) (*domain.User, error) {

	return persistence.NewUserRepositoryWithRDB().Get(id)
}

func UsersUpdate(id int, name string) (*domain.User, error) {

	user, _ := UsersGet(id)
	user.Name = name
	return persistence.NewUserRepositoryWithRDB().Update(user)
}
