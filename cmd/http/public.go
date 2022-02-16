package http

import (
	"basic_golang/config"
	"basic_golang/internal/adapter"
	"context"

	"basic_golang/internal/handler/public"

	"github.com/spf13/cobra"
)

var PublicCmd = &cobra.Command{
	Use:   "http-public",
	Short: "Starts Public REST API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.LoadMainConfig()

		database, _ := adapter.NewSqliteAdapter(context.TODO())

		srv := public.NewServer(cfg, database)
		return srv.Run()
	},
}
