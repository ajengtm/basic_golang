package services

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/auth/entity"
	"context"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (s *authDomain) CheckToken(ctx context.Context, jwtToken string) (user entity.User, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Auth_CheckToken")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(jwtToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			logger.Error("error when CheckToken", zap.Error(err))
			return user, fmt.Errorf("Not Authorized")
		}
		return
	}

	if !tkn.Valid {
		logger.Error("error when CheckToken", zap.Error(err))
		return user, fmt.Errorf("Not Authorized")
	}

	return entity.User{
		Username:  claims.Username,
		Phone:     claims.Phone,
		Role:      claims.Role,
		Timestamp: claims.Timestamp,
	}, nil
}
