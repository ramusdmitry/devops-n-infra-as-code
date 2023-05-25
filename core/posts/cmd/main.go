package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	postsApp "posts-app-service"
	"posts-app-service/pkg/handler"
	"posts-app-service/pkg/probes"
	"posts-app-service/pkg/repository"
	"posts-app-service/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Info("Posts-app-service is starting...")

	probe := probes.NewProbes()

	cfg, err := postsApp.LoadConfig("configs", "config")

	if err != nil {
		logrus.Fatalf("Failed to init config/envs: %s", err.Error())
	}

	probErr := probe.InitProbes(cfg)

	if probErr != nil {
		logrus.Fatalf("[probes] service started failed %s", probErr.Error())
	}

	db, err := repository.NewPostgresDB(cfg.DB)

	if err != nil {
		logrus.Fatalf("Failed to init database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(postsApp.Server)
	if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error", err.Error())
	}

}
