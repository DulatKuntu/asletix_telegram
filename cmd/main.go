package main

import (
	"asletix_telegram"
	"asletix_telegram/pkg/handler"
	"asletix_telegram/pkg/repository"
	"asletix_telegram/pkg/service"
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, please set")
	}
}

func main() {
	if err := initEnv(); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	db, client, err := repository.NewMongoDB(
		repository.Config{
			DBURI:  os.Getenv("MONGO_HOST"),
			DBName: "asletix",
		},
	)

	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(asletix_telegram.Server)

	if err := srv.RunTLS(os.Getenv("WebhookPort"), os.Getenv("LocationCertificate")+"asletix_com.pem", os.Getenv("LocationCertificate")+"asletix.key", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initEnv() error {
	_, exists := os.LookupEnv("MONGO_HOST")

	if !exists {
		os.Setenv("MONGO_HOST", "mongodb://localhost:27017")
	}

	reqs := []string{
		"WebhookPort",
		"LocationCertificate",
	}

	for i := 0; i < len(reqs); i++ {
		_, exists = os.LookupEnv(reqs[i])

		if !exists {
			return errors.New(".env variables not set")
		}
	}

	return nil
}
