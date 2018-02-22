package repository

import "github.com/wmetaw/go-ddd-on-echo/domain"

// Domain Service層

// infrastructure層の「infrastructure/persistance/user_repository」が
// 実装するメソッド(インターフェース) を定義する
// これにより、モデルがインフラに依存しなくなり、依存性の逆転が生まれる（インフラ側の実装を別のDBに差し替え易い）
type UserRepository interface {
	Get(id int) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(*domain.User) (*domain.User, error)
}
