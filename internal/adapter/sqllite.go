package adapter

import (
	zaplogger "basic_golang/internal/adapter/zap"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/opentracing/opentracing-go"
)

func NewSqliteAdapter(ctx context.Context) (database *sql.DB, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "adapter/NewSqliteAdapter")
	defer span.Finish()
	logger := zaplogger.For(ctx)

	database, err = sql.Open("sqlite3", "./basic_golang.db")
	if err != nil {
		return database, err
	}
	logger.Info("sqlite connected")

	statement, err := database.Prepare(`
		CREATE TABLE IF NOT EXISTS user 
		(id INTEGER PRIMARY KEY, 
			username TEXT, 
			phone TEXT,
			role TEXT,
			password TEXT,
			created_at TEXT)`)
	statement.Exec()
	if err != nil {
		return database, err
	}

	return database, nil
}
