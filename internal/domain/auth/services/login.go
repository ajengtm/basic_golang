package services

import (
	"context"
	"fmt"

	zaplogger "basic_golang/internal/adapter/zap"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"6281381209988": "password1",
	"6285771247280": "password2",
}

type LoginRequest struct {
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

type Claims struct {
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
	jwt.StandardClaims
}

func (s *authDomain) Login(ctx context.Context, inputLogin *LoginRequest) (jwtToken string, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Auth_Login")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	user, err := s.authRepository.Find(ctx, "phone", inputLogin.Phone)
	if err != nil {
		logger.Error("error when Login|FindByPhoneNumber", zap.Error(err))
		return jwtToken, err
	}

	if user.Password != inputLogin.Password {
		logger.Error("error when Login, password not match", zap.Error(err))
		return jwtToken, fmt.Errorf("Not Authorized")
	}

	claims := &Claims{
		Username:  user.Username,
		Phone:     user.Phone,
		Role:      user.Role,
		Timestamp: user.Timestamp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.Error("error when Login generate token", zap.Error(err))
		return jwtToken, err
	}

	return tokenString, nil
}
