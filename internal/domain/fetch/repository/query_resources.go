package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/fetch/entity"
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) CountResource(ctx context.Context) (count int, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository_Fetch_CountResource")
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

func (r *fetchRepository) FindResources(ctx context.Context) (resources []entity.Resource, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository_Fetch_GetResourcesAgregation")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	stmt, err := r.database.Prepare(`
		SELECT uuid, komoditas, area_provinsi, area_kota, size, price, tgl_parsed, timestamp
		FROM resource `)
	rows, err := stmt.Query()
	if err != nil {
		zaplogger.For(ctx).Error("Database error", zap.Error(err))
		return resources, err
	}
	defer rows.Close()
	for rows.Next() {
		resource := entity.Resource{}
		if err = rows.Scan(
			&resource.UUID,
			&resource.Komoditas,
			&resource.AreaProvinsi,
			&resource.AreaKota,
			&resource.Size,
			&resource.Price,
			&resource.ParsedDate,
			&resource.Timestamp,
		); err != nil {
			logger.Error("Database error", zap.Error(err))
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}

func (r *fetchRepository) GetResourcesAgregation(ctx context.Context, functions string) (resources []entity.AggregateResources, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository_Fetch_GetResourcesAgregation")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	stmt, err := r.database.Prepare(`
	SELECT area_provinsi 
	,` + functions + `(tgl_parsed) as week_begining
	FROM resource 
	GROUP BY area_provinsi	`)
	rows, err := stmt.Query()
	if err != nil {
		zaplogger.For(ctx).Error("Database error", zap.Error(err))
		return resources, err
	}
	defer rows.Close()
	for rows.Next() {
		resource := entity.AggregateResources{}
		if err = rows.Scan(
			&resource.AreaProvinsi,
			&resource.WeekBegining,
		); err != nil {
			logger.Error("Database error", zap.Error(err))
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}
