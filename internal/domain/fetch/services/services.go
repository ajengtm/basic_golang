package services

import (
	"basic_golang/internal/domain/auth"
	"basic_golang/internal/domain/fetch/repository"
)

type fetchDomain struct {
	authDomain      auth.AuthDomainInterface
	fetchRepository repository.FetchRepositoryInterface
}

func NewFetchServices(authDomain auth.AuthDomainInterface, fetchRepository repository.FetchRepositoryInterface) *fetchDomain {
	return &fetchDomain{
		authDomain,
		fetchRepository,
	}
}
