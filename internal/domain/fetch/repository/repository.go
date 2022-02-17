package repository

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"basic_golang/internal/domain/fetch/entity"
	"context"

	_ "github.com/mattn/go-sqlite3"
)

type fetchRepository struct {
	cfg     config.MainConfig
	myCache adapter.CacheItf
}

type FetchRepositoryInterface interface {
	GetResources(ctx context.Context) (resources []entity.Resource, err error)
	GetCurrencyIDRtoUSD(ctx context.Context) (entity.UsdIdrCurrency, error)
}

func NewFetchRepository(cfg config.MainConfig, myCache adapter.CacheItf) FetchRepositoryInterface {
	fetchRepo := &fetchRepository{
		cfg:     cfg,
		myCache: myCache,
	}
	return fetchRepo
}
