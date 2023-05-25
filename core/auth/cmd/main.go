package main

import (
	authApp "auth-app-service"
	"auth-app-service/pkg/handler"
	metrics "auth-app-service/pkg/metrics"
	"auth-app-service/pkg/probes"
	"auth-app-service/pkg/repository"
	"auth-app-service/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Info("Auth-app-service is starting...")

	probe := probes.NewProbes()

	cfg, err := authApp.LoadConfig("configs", "config")

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

	metricsCollector := metrics.NewMetrics()
	go func() {
		if err = metricsCollector.Serve(cfg); err != nil {
			logrus.Fatalf("failed to start metrics collector: %s", err.Error())
		}
	}()

	repos := repository.NewRepository(db)
	services := service.NewService(repos, metricsCollector)
	handlers := handler.NewHandler(services)

	server := new(authApp.Server)
	if err := server.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error with running server: %s", err.Error())
	}
}
