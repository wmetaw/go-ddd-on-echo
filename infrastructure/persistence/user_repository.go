package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wmetaw/go-ddd-on-echo/domain"
	"github.com/wmetaw/go-ddd-on-echo/domain/repository"
)

// UserRepository Implements repository.UserRepository
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepositoryWithRDB returns initialized UserRepositoryImpl
// 戻り値をinterfaceにすることでUserRepository interfaceを全て実装しないとエラー
func NewUserRepositoryWithRDB(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetAll() ([]*domain.User, error) {

	// usersインスタンス化
	users := []*domain.User{}
	r.DB.Find(&users)

	return users, nil
}

func (r *UserRepositoryImpl) Get(id int) (*domain.User, error) {

	if id < 1 {
		panic("no value for param requested")
	}

	user := domain.User{}
	r.DB.Find(&user, id)

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {

	// transaction
	tx := r.DB.Begin()
	fmt.Println(user)
	if err := tx.Model(&user).UpdateColumn("name", user.Name).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return user, nil
}
