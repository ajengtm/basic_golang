package fetch

import (
	"basic_golang/internal/domain/auth"
	"basic_golang/internal/domain/fetch/entity"
	"basic_golang/internal/domain/fetch/services"
	"context"
)

type FetchDomainInterface interface {
	GetResourcesAdmin(ctx context.Context, jwtToken string) ([]entity.Resource, error)
	GetResources(ctx context.Context, jwtToken string) ([]entity.Resource, error)
}

func NewFetchDomain(authDomain auth.AuthDomainInterface) FetchDomainInterface {
	fetchServices := services.NewFetchServices(authDomain)
	return fetchServices
}
