package adapter

import (
	"context"

	"basic_golang/config"
	zaplogger "basic_golang/internal/adapter/zap"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMysqlAdapter connection init func
func NewMysqlAdapter(ctx context.Context, config config.MainConfig) (dbMaster *gorm.DB, dbSlave *gorm.DB) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "adapter/mysql/NewMysqlAdapter")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	dbMaster = NewGormDB(ctx, mysql.Config{
		DSN: config.DBConfig.MasterDSN,
	})
	if !config.DBConfig.MasterOnly {
		dbSlave = NewGormDB(ctx, mysql.Config{
			DSN: config.DBConfig.SlaveDSN,
		})
	}
	logger.Info("mysql connected")

	if config.DBConfig.DebugMode {
		dbMaster = dbMaster.Debug()
		if dbSlave != nil {
			dbSlave = dbSlave.Debug()
		}
	}

	return dbMaster, dbSlave
}

// NewGormDB create db connection
func NewGormDB(ctx context.Context, config mysql.Config) *gorm.DB {
	span, ctx := opentracing.StartSpanFromContext(ctx, "adapter/mysql/NewGormDB")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	db, err := gorm.Open(mysql.New(config), &gorm.Config{})
	if err != nil {
		logger.Panic("error when try to open new mysql conn", zap.Error(err))
	}

	return db
}
