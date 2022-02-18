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
	// related to stein.efishery.com
	GetResources(ctx context.Context) (resources []entity.Resource, err error)

	// related to free.currencyconverterapi.com
	GetCurrencyIDRtoUSD(ctx context.Context) (entity.UsdIdrCurrency, error)
	GetCurrencyFromCache(ctx context.Context) (currency float64)
	GetCurrencyFromCurrconv(ctx context.Context) (currency entity.UsdIdrCurrency, err error)
}

func NewFetchRepository(cfg config.MainConfig, myCache adapter.CacheItf) FetchRepositoryInterface {
	fetchRepo := &fetchRepository{
		cfg:     cfg,
		myCache: myCache,
	}
	return fetchRepo
}
