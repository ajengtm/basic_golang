package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) CountResource(ctx context.Context) (count int, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mysql_Fetch_FindByUUID")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	stmt, err := r.database.Prepare(`
		SELECT count(uuid)
		FROM resource `)
	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when CountResource to database", zap.Error(err))
		return count, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			logger.Error("error when CountResource to database", zap.Error(err))
			return count, err
		}

	}

	return count, nil
}
