package main

import (
	"log"
	"os"

	http "basic_golang/cmd/http"
	zaplogger "basic_golang/internal/adapter/zap"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "basic-go",
	}

	// init zap logger
	zaplogger.Init()

	// Add Command
	rootCmd.AddCommand(http.PublicCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
