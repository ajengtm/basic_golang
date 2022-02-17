package fetch

import (
	"basic_golang/config"
	"basic_golang/internal/domain/auth"
	"basic_golang/internal/domain/fetch/entity"
	"basic_golang/internal/domain/fetch/repository"
	"basic_golang/internal/domain/fetch/services"
	"context"
)

type FetchDomainInterface interface {
	GetResourcesAdmin(ctx context.Context, jwtToken string) ([]entity.Resource, error)
	GetResources(ctx context.Context, jwtToken string) ([]entity.ResourceResponse, error)
}

func NewFetchDomain(cfg config.MainConfig, authDomain auth.AuthDomainInterface) FetchDomainInterface {
	fetchRepo := repository.NewFetchRepository(cfg)
	fetchServices := services.NewFetchServices(authDomain, fetchRepo)
	return fetchServices
}
