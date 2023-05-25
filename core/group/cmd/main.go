package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	groupApp "group-app-service"
	"group-app-service/pkg/handler"
	"group-app-service/pkg/probes"
	"group-app-service/pkg/repository"
	"group-app-service/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Info("Group-app-service starting...")

	probe := probes.NewProbes()

	cfg, err := groupApp.LoadConfig("configs", "config")

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

	server := new(groupApp.Server)
	if err := server.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error with running server: %s", err.Error())
	}

}
