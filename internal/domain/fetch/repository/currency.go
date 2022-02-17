package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/fetch/entity"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) GetCurrencyIDRtoUSD(ctx context.Context) (currency entity.UsdIdrCurrency, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Fetch_Getcurrency")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, r.cfg.CurrencyConverter.URL+r.cfg.CurrencyConverter.APIkey, nil)
	if err != nil {
		logger.Error("REPO||FETCH||Getcurrency||failed to create new request", zap.Error(err))
		return currency, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("REPO||FETCH||Getcurrency||error when try to request to free.currconv.com", zap.Error(err))
		return currency, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("REPO||FETCH||Getcurrency||error when try to read response body from free.currconv.com", zap.Error(err))
		return currency, err
	}

	logger.Info(
		"REPO||FETCH||Getcurrency||Response from free.currconv.com",
		zap.String("responseBody", string(respBody)),
		zap.Int("statusCode", resp.StatusCode),
	)

	if resp.StatusCode != http.StatusOK {
		logger.Error(
			"REPO||FETCH||Getcurrency||error response from free.currconv.com",
			zap.Int("StatusCode", resp.StatusCode),
			zap.String("RespBody", string(respBody)),
		)
		return currency, errors.New(string(respBody))
	}

	if err := json.Unmarshal([]byte(respBody), &currency); err != nil {
		logger.Error("REPO||FETCH||Getcurrency||error when try to decode response body from free.currconv.com", zap.Error(err))
		return currency, err
	}

	logger.Info("REPO||FETCH||Getcurrency||Success")
	return currency, nil
}
