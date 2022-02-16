package http

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"context"

	"basic_golang/internal/handler/public_fetch"

	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:   "http-fetch",
	Short: "Starts FETCH REST API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.LoadMainConfig()

		database, _ := adapter.NewSqliteAdapter(context.TODO())

		srv := public_fetch.NewServer(cfg, database)
		return srv.Run()
	},
}
