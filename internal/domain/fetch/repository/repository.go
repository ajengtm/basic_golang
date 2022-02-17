package repository

import (
	"basic_golang/config"
	"basic_golang/internal/domain/fetch/entity"
	"context"

	_ "github.com/mattn/go-sqlite3"
)

type fetchRepository struct {
	cfg config.MainConfig
}

type FetchRepositoryInterface interface {
	GetResources(ctx context.Context) (resources []entity.Resource, err error)
	GetCurrencyIDRtoUSD(ctx context.Context) (entity.UsdIdrCurrency, error)
}

func NewFetchRepository(cfg config.MainConfig) FetchRepositoryInterface {
	fetchRepo := &fetchRepository{
		cfg: cfg,
	}
	return fetchRepo
}
