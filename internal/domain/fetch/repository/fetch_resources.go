package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/fetch/entity"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) GetResources(ctx context.Context) (resources []entity.Resource, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Fetch_GetResources")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, r.cfg.Efishery.URL, nil)
	if err != nil {
		logger.Error("REPO||FETCH||GetResources||failed to create new request", zap.Error(err))
		return resources, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("REPO||FETCH||GetResources||error when try to request to stein.efishery.com", zap.Error(err))
		return resources, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("REPO||FETCH||GetResources||error when try to read response body from stein.efishery.com", zap.Error(err))
		return resources, err
	}

	logger.Info(
		"REPO||FETCH||GetResources||Response from stein.efishery.com",
		zap.String("responseBody", string(respBody)),
		zap.Int("statusCode", resp.StatusCode),
	)

	if resp.StatusCode != http.StatusOK {
		logger.Error(
			"REPO||FETCH||GetResources||error response from stein.efishery.com",
			zap.Int("StatusCode", resp.StatusCode),
			zap.String("RespBody", string(respBody)),
		)
		return resources, errors.New(string(respBody))
	}

	if err := json.Unmarshal([]byte(respBody), &resources); err != nil {
		logger.Error("REPO||FETCH||GetResources||error when try to decode response body from stein.efishery.com", zap.Error(err))
		return resources, err
	}

	logger.Info("REPO||FETCH||GetResources||Success")
	return
}
