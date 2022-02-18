package repository

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/fetch/entity"
	"context"
	"strings"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (r *fetchRepository) InsertResources(ctx context.Context, resources []entity.Resource) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mysql_Fetch_InsertResources")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	sqlStr := `INSERT INTO resource (uuid, komoditas, area_provinsi, area_kota, size, price, tgl_parsed, timestamp) VALUES `

	vals := []interface{}{}
	for _, resource := range resources {
		sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals,
			resource.UUID,
			resource.Komoditas,
			resource.AreaProvinsi,
			resource.AreaKota,
			resource.Size,
			resource.Price,
			resource.ParsedDate,
			resource.Timestamp,
		)

	}
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	_, err = r.database.Exec(sqlStr, vals...)
	if err != nil {
		logger.Error("Database error", zap.Error(err))
		return err
	}
	return nil
}
