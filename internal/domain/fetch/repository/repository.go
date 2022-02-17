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
	InsertResources(ctx context.Context, resources []entity.Resource) (err error)
	CountResource(ctx context.Context) (count int, err error)
	GetResources(ctx context.Context) (resources []entity.Resource, err error)
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
