package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/zelta-7/img-server/pkg/config"
	"github.com/zelta-7/img-server/pkg/repository"
	"github.com/zelta-7/img-server/pkg/service"
	"github.com/zelta-7/img-server/pkg/transport"
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

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Name, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name)
	fmt.Println(dbConnectionString)
	db, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&repository.ImageRecord{})
	imgRepo := repository.NewImgRepo(db)
	log.Print("Repository layer initialised.")

	imgService := service.NewImgService(imgRepo)
	log.Print("Service layer initialised.")

	r, err := transport.Route(imgService)
	if err != nil {
		log.Print("Error creating router")
		panic(err)
	}

	r.Run(":8080")
}

func run(cfg *config.Config) {
	cfg.SetConfigFilePath().LoadConfigFile()
	cfg.InitImgServer()

	if !cfg.Enabled {
		log.Fatalf("Server is not enabled")
	}
	log.Printf("Server listening on port: %s", cfg.Port)

}
