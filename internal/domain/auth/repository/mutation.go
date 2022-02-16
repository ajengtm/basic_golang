package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/auth/entity"
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (c *authRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mysql_Auth_CreateUser")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	stmt, err := c.Database.Prepare(`
	INSERT INTO 
		user(phone,role,username,password,created_at) 
	VALUES (?,?,?,?,?)`)
	if err != nil {
		logger.Error("error when CreateUser to database", zap.Error(err))
		return user, err
	}

	result, err := stmt.Exec(user.Phone, user.Role, user.Username, user.Password, user.Timestamp)
	if err != nil {
		logger.Error("error when CreateUser to database", zap.Error(err))
		return user, err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when CreateUser to database", zap.Error(err))
		return user, err
	}
	user.ID = newID

	return user, nil
}
