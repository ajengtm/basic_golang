package http

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"basic_golang/internal/handler/public_auth"
	"context"

	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "http-auth",
	Short: "Starts AUTH REST API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.LoadMainConfig()

		database, _ := adapter.NewSqliteAdapter(context.TODO())

		srv := public_auth.NewServer(cfg, database)
		return srv.Run()
	},
}
