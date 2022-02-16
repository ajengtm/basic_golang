package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/auth/entity"
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (c *authRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (password string, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mysql_Auth_FindByPhoneNumber")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	stmt, err := c.Database.Prepare(`
		SELECT password
		FROM user 
		where phone = ?
		Limit 1`)
	rows, err := stmt.Query(phoneNumber)
	if err != nil {
		logger.Error("error when FindByPhoneNumber to database", zap.Error(err))
		return password, err
	}

	for rows.Next() {
		err = rows.Scan(&password)
		if err != nil {
			logger.Error("error when FindByPhoneNumber to database", zap.Error(err))
			return password, err
		}

	}

	return password, nil
}

func (c *authRepository) Find(ctx context.Context, filterBy string, filterValue string) (user entity.User, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mysql_Auth_FindByUsername")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	stmt, err := c.Database.Prepare(`
		SELECT id, username, phone, role, password, created_at
		FROM user 
		where ` + filterBy + ` = ?
		Limit 1`)

	rows, err := stmt.Query(filterValue)
	if err != nil {
		logger.Error("error when FindUserByUsername to database", zap.Error(err))
		return user, err
	}
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Phone,
			&user.Role,
			&user.Password,
			&user.Timestamp,
		)
		if err != nil {
			logger.Error("error when FindUserByUsername to database", zap.Error(err))
		}
	}

	return user, nil
}
