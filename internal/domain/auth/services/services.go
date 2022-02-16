package services

import (
	"basic_golang/internal/domain/auth/repository"
)

type authDomain struct {
	authRepository repository.AuthRepositoryInterface
}

func NewAuthServices(authRepository repository.AuthRepositoryInterface) *authDomain {
	return &authDomain{
		authRepository,
	}
}
