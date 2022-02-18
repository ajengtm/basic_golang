package repository

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"basic_golang/internal/domain/fetch/entity"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type fetchRepository struct {
	cfg      config.MainConfig
	database *sql.DB
	myCache  adapter.CacheItf
}

type FetchRepositoryInterface interface {
	// related to stein.efishery.com
	GetResources(ctx context.Context) (resources []entity.Resource, err error)

	// related to db
	InsertResources(ctx context.Context, resources []entity.Resource) (err error)
	CountResource(ctx context.Context) (count int, err error)
	FindResources(ctx context.Context) (resources []entity.Resource, err error)
	GetResourcesAgregation(ctx context.Context, functions string) (resources []entity.AggregateResources, err error)

	// related to free.currencyconverterapi.com
	GetCurrencyIDRtoUSD(ctx context.Context) (entity.UsdIdrCurrency, error)
	GetCurrencyFromCache(ctx context.Context) (currency float64)
	GetCurrencyFromCurrconv(ctx context.Context) (currency entity.UsdIdrCurrency, err error)
}

func NewFetchRepository(cfg config.MainConfig, database *sql.DB, myCache adapter.CacheItf) FetchRepositoryInterface {
	fetchRepo := &fetchRepository{
		cfg:      cfg,
		myCache:  myCache,
		database: database,
	}
	return fetchRepo
}
