package fetch

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"basic_golang/internal/domain/auth"
	"basic_golang/internal/domain/fetch/entity"
	"basic_golang/internal/domain/fetch/repository"
	"basic_golang/internal/domain/fetch/services"
	"context"
	"database/sql"
)

type FetchDomainInterface interface {
	GetResourcesAdmin(ctx context.Context, jwtToken string) (res entity.ResourceAgregationResponse, err error)
	GetResources(ctx context.Context, jwtToken string) ([]entity.ResourceResponse, error)
	SeedDataResources(ctx context.Context) (res []entity.Resource, err error)
}

func NewFetchDomain(cfg config.MainConfig, database *sql.DB, myCache adapter.CacheItf, authDomain auth.AuthDomainInterface) FetchDomainInterface {
	fetchRepo := repository.NewFetchRepository(cfg, database, myCache)
	fetchServices := services.NewFetchServices(authDomain, fetchRepo)
	return fetchServices
}
