package repository

import (
	"todo-rest-api-3/model"

	"gorm.io/gorm"
)

// インターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// インターフェースを実装する構造体
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// 実装部
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
