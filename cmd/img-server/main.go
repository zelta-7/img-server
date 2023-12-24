package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zelta-7/img-server/pkg/config"
)

func main() {
	cfg := config.Default()

	cmd := &cobra.Command{
		Use:   "img-server",
		Short: "process images",
		Long:  "detailed process images",
		Run: func(cmd *cobra.Command, args []string) {
			run(cfg)
		},
	}

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Error in executing the server: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(cfg *config.Config) {
	cfg.SetConfigFilePath().LoadConfigFile()

	if !cfg.Enabled {
		log.Fatalf("Server is not enabled")
	}
	cfg.InitImgServer()
	log.Printf("Server listening on port: %s", cfg.Port)

}
