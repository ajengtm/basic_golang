package repository

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type authRepository struct {
	// Database *gorm.DB
}

type AuthRepositoryInterface interface {
}

func NewAuthRepository() AuthRepositoryInterface {
	authRepo := &authRepository{
		// Database: gormdb,
	}
	return authRepo
}
