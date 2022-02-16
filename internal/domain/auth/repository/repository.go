package repository

import (
	"basic_golang/internal/domain/auth/entity"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type authRepository struct {
	Database *sql.DB
}

type AuthRepositoryInterface interface {
	Find(ctx context.Context, filterBy string, filterValue string) (user entity.User, err error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}

func NewAuthRepository(database *sql.DB) AuthRepositoryInterface {
	authRepo := &authRepository{
		Database: database,
	}
	return authRepo
}
