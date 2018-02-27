package persistence

import (
	"fmt"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"github.com/wmetaw/go-ddd-on-echo/domain"
	"github.com/wmetaw/go-ddd-on-echo/domain/repository"
)

// UserRepository Implements repository.UserRepository
type UserRepositoryImpl struct{}

// NewUserRepositoryWithRDB returns initialized UserRepositoryImpl
// 戻り値をinterfaceにすることでUserRepository interfaceを全て実装しないとエラー
func NewUserRepositoryWithRDB() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) GetAll() ([]*domain.User, error) {

	// usersインスタンス化
	users := []*domain.User{}
	config.DBCon.Find(&users)

	return users, nil
}

func (r *UserRepositoryImpl) Get(id int) (*domain.User, error) {

	user := domain.User{}
	config.DBCon.Find(&user, id)

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {

	// transaction
	tx := config.DBCon.Begin()
	fmt.Println(user)
	if err := tx.Model(&user).UpdateColumn("name", user.Name).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return user, nil
}
