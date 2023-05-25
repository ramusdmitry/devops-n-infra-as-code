package main

import (
	commsApp "comms-app-service"
	"comms-app-service/pkg/handler"
	"comms-app-service/pkg/probes"
	"comms-app-service/pkg/repository"
	"comms-app-service/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Infof("Comms-app-service is starting...")

	probe := probes.NewProbes()

	cfg, err := commsApp.LoadConfig("configs", "config")

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

	srv := new(commsApp.Server)
	if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error", err.Error())
	}

}
