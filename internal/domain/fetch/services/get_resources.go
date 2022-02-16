package services

import (
	"context"
	"fmt"

	zaplogger "basic_golang/internal/adapter/zap"
	authEntity "basic_golang/internal/domain/auth/entity"
	"basic_golang/internal/domain/fetch/entity"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (s *fetchDomain) GetResources(ctx context.Context, jwtToken string) (res []entity.Resource, err error) {
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
		return res, fmt.Errorf("Not Authorized")
	}

	return res, nil
}
