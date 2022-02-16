package auth

import (
	"basic_golang/internal/domain/auth/entity"
	"basic_golang/internal/domain/auth/services"
	"context"
	"database/sql"

	"basic_golang/internal/domain/auth/repository"
)

type AuthDomainInterface interface {
	Login(ctx context.Context, inputLogin *services.LoginRequest) (jwtToken string, err error)
	CheckToken(ctx context.Context, jwtToken string) (user entity.User, err error)
	UpsertUser(ctx context.Context, inputUser *services.UserRequest) (user entity.User, err error)
}

func NewAuthDomain(database *sql.DB) AuthDomainInterface {
	authRepo := repository.NewAuthRepository(database)
	authServices := services.NewAuthServices(authRepo)
	return authServices
}
