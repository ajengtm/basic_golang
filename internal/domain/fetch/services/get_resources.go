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

func (s *fetchDomain) GetResourcesAdmin(ctx context.Context, jwtToken string) (res entity.ResourceAgregationResponse, err error) {
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

	count, err := s.fetchRepository.CountResource(ctx)
	if err != nil {
		logger.Error("error when GetResources|CountResource", zap.Error(err))
		return res, err
	}

	if count == 0 {
		_, err = s.SeedDataResources(ctx)
		if err != nil {
			logger.Error("error when GetResources|seedDataResources", zap.Error(err))
			return res, err
		}
	}

	sqlFunction := []string{"MIN", "MAX", "AVG"}
	for _, v := range sqlFunction {
		aggregateResources, err := s.fetchRepository.GetResourcesAgregation(ctx, v)
		if err != nil {
			logger.Error("error when GetResources|GetResourcesAgregation", zap.Error(err))
			return res, err
		}
		switch v {
		case "MIN":
			res.Min = aggregateResources
		case "MAX":
			res.Max = aggregateResources
		case "AVG":
			res.Avg = aggregateResources
		}
	}

	return res, nil
}

func (s *fetchDomain) SeedDataResources(ctx context.Context) (res []entity.Resource, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Fetch_seedDataResources")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	resources, err := s.fetchRepository.GetResources(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}
	var newResources []entity.Resource
	for _, resource := range resources {
		if (resource != entity.Resource{}) || resource.UUID != "" {
			newResources = append(newResources, resource)
		}
	}

	err = s.fetchRepository.InsertResources(ctx, newResources)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}

	return newResources, nil
}
