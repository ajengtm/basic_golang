package services

import (
	"context"
	"fmt"
	"strconv"

	zaplogger "basic_golang/internal/adapter/zap"
	authEntity "basic_golang/internal/domain/auth/entity"
	"basic_golang/internal/domain/fetch/entity"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (s *fetchDomain) GetResources(ctx context.Context, jwtToken string) (res []entity.ResourceResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Fetch_GetResources")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	user, err := s.authDomain.CheckToken(ctx, jwtToken)
	if err != nil {
		logger.Error("error when GetResources|CheckToken", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}
	if (user == authEntity.User{}) {
		return res, fmt.Errorf("Not Authorized")
	}

	resources, err := s.fetchRepository.GetResources(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}

	usdIdrCurrency, err := s.fetchRepository.GetCurrencyIDRtoUSD(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetCurrencyUSDToIDR", zap.Error(err))
		return res, err
	}

	for _, resource := range resources {
		if (resource != entity.Resource{}) {
			var price float64
			if resource.Price != "" {
				price, err = strconv.ParseFloat(resource.Price, 32)
				if err != nil {
					logger.Error("error when GetResources|ParseFloat Price", zap.Error(err))
				}
			}

			res = append(res, entity.ResourceResponse{
				UUID:         resource.UUID,
				Komoditas:    resource.Komoditas,
				AreaProvinsi: resource.Komoditas,
				AreaKota:     resource.AreaKota,
				Size:         resource.Size,
				Price:        price,
				ParsedDate:   resource.ParsedDate,
				Timestamp:    resource.Timestamp,
				PriceUSD:     price * usdIdrCurrency.IDRtoUSD,
			})
		}
	}

	return res, nil
}

func (s *fetchDomain) GetResourcesAdmin(ctx context.Context, jwtToken string) (res []entity.Resource, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Fetch_GetResources")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	user, err := s.authDomain.CheckToken(ctx, jwtToken)
	if err != nil {
		logger.Error("error when GetResources|CheckToken", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}

	if (user == authEntity.User{}) {
		logger.Error("error when GetResources|CheckToken User not found", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}
	if user.Role != "admin" {
		logger.Error("error when GetResources|CheckToken Invalid User Role", zap.Error(err))
		return res, fmt.Errorf("Not Authorized - Invalid User Role")
	}

	resources, err := s.fetchRepository.GetResources(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}

	for _, resource := range resources {
		if (resource != entity.Resource{}) {
			res = append(res, resource)
		}
	}

	return res, nil
}
