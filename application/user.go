package application

import (
	"github.com/wmetaw/go-ddd-on-echo/config"
	"github.com/wmetaw/go-ddd-on-echo/domain"
	"github.com/wmetaw/go-ddd-on-echo/infrastructure/persistence"
)

// GetUsers returns user list
func Users() ([]*domain.User, error) {
	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := persistence.NewUserRepositoryWithRDB(db)
	return repo.GetAll()
}

// GetUser returns user list
func UsersGet(id int) (*domain.User, error) {
	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := persistence.NewUserRepositoryWithRDB(db)
	return repo.Get(id)
}

func UsersUpdate(id int, name string) (*domain.User, error) {

	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, _ := UsersGet(id)
	user.Name = name

	repo := persistence.NewUserRepositoryWithRDB(db)
	return repo.Update(user)
}
