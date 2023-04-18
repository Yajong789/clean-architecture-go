package users

import (
	"gorm.io/gorm"
)

type Repository struct{
	DB *gorm.DB
}

func (repo Repository) CheckUsername(username string) (*User, error) {
	var user User
	result := repo.DB.Where("username", username).First(&user)

	return &user, result.Error
}

func (repo Repository) CreateUser(user *User) error{
	result := repo.DB.Create(&user)

	return result.Error
}
