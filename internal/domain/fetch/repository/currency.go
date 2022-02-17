package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/fetch/entity"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) GetCurrencyIDRtoUSD(ctx context.Context) (entity.UsdIdrCurrency, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Fetch_Getcurrency")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	cacheCurrency := r.GetCurrencyFromCache(ctx)
	logger.Info(
		"REPO||FETCH||GetCurrencyIDRtoUSD||GetCurrencyFromCache",
		zap.Float64("cacheCurrency", cacheCurrency),
	)
	if cacheCurrency != 0 {
		return entity.UsdIdrCurrency{
			IDRtoUSD: cacheCurrency,
		}, nil
	} else {
		currconvCurrency, err := r.GetCurrencyFromCurrconv(ctx)
		if err != nil {
			logger.Error("REPO||FETCH||GetCurrencyIDRtoUSD||error when get currency from free.currconv.com", zap.Error(err))
			return entity.UsdIdrCurrency{}, err
		}
		return currconvCurrency, nil
	}
}

func (r *fetchRepository) GetCurrencyFromCache(ctx context.Context) (currency float64) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Fetch_GetCurrencyFromCache")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	cacheCurrency, err := r.myCache.Get("IDRtoUSD")
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCache||error when get currency from local cache", zap.Error(err))
		return 0
	}
	if cacheCurrency == nil {
		return 0
	}

	err = json.Unmarshal(cacheCurrency, &currency)
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCache||error when get currency from local cache", zap.Error(err))
		return 0
	}
	logger.Info("REPO||FETCH||GetCurrencyFromCache||Success cache exist")
	return currency
}

func (r *fetchRepository) GetCurrencyFromCurrconv(ctx context.Context) (currency entity.UsdIdrCurrency, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Fetch_GetCurrencyFromCurrconv")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, r.cfg.CurrencyConverter.URL+r.cfg.CurrencyConverter.APIkey, nil)
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCurrconv||failed to create new request", zap.Error(err))
		return currency, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCurrconv||error when try to request to free.currconv.com", zap.Error(err))
		return currency, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCurrconv||error when try to read response body from free.currconv.com", zap.Error(err))
		return currency, err
	}

	logger.Info(
		"REPO||FETCH||GetCurrencyFromCurrconv||Response from free.currconv.com",
		zap.String("responseBody", string(respBody)),
		zap.Int("statusCode", resp.StatusCode),
	)

	if resp.StatusCode != http.StatusOK {
		logger.Error(
			"REPO||FETCH||GetCurrencyFromCurrconv||error response from free.currconv.com",
			zap.Int("StatusCode", resp.StatusCode),
			zap.String("RespBody", string(respBody)),
		)
		return currency, errors.New(string(respBody))
	}

	if err := json.Unmarshal([]byte(respBody), &currency); err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCurrconv||error when try to decode response body from free.currconv.com", zap.Error(err))
		return currency, err
	}

	err = r.myCache.Set("IDRtoUSD", currency.IDRtoUSD, 10*time.Minute)
	if err != nil {
		logger.Error("REPO||FETCH||GetCurrencyFromCurrconv||error when try SET currency to cache", zap.Error(err))
		return currency, err
	}

	logger.Info("REPO||FETCH||GetCurrencyFromCurrconv||Success")

	return currency, nil
}
