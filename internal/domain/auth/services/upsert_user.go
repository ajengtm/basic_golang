package services

import (
	"basic_golang/internal/adapter"
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/auth/entity"
	"context"
	"math/rand"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type UserRequest struct {
	Username string `json:"username,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (s *authDomain) UpsertUser(ctx context.Context, inputUser *UserRequest) (user entity.User, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_UpsertUser")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	user, err = s.authRepository.Find(ctx, "username", inputUser.Username)
	if err != nil {
		logger.Error("error when UpsertUser|FindByUsername", zap.Error(err))
		return user, err
	}
	if (entity.User{} != user) {
		return user, nil
	}

	// tidak ketemu insert user baru ke db
	user.Phone = inputUser.Phone
	user.Role = inputUser.Role
	user.Username = inputUser.Username
	user.Password = generatePassword()
	user.Timestamp = adapter.GetCurrentTimestampTZ()
	user, err = s.authRepository.CreateUser(ctx, user)
	if err != nil {
		logger.Error("error when UpsertUser|FindByUsername", zap.Error(err))
		return user, err
	}

	return user, nil
}

func generatePassword() string {
	const passwordLen = 4
	passwordRunes := make([]rune, passwordLen)

	for i := range passwordRunes {
		passwordRunes[i] = randomRune()
	}
	return string(passwordRunes)
}

func randomRune() rune {
	rand.Seed(time.Now().UnixNano())

	f := rand.Float64()
	var min, max rune

	switch {
	case f < float64(1)/3:
		min, max = 'a', 'z'
	case f < float64(2)/3:
		min, max = 'A', 'Z'
	default:
		min, max = '0', '9'
	}

	delta := int(max - min)
	i := rand.Intn(delta + 1)

	return min + rune(i)
}
