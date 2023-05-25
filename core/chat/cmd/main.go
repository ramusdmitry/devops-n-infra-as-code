package main

import (
	chatApp "chat-app-service"
	"chat-app-service/pkg/service"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Info("Chat-app-service starting...")

	cfg, err := chatApp.LoadConfig("configs", "config")

	if err != nil {
		logrus.Fatalf("Failed to init config/envs: %s", err.Error())
	}

	//probe := probes.NewProbes()
	//probErr := probe.InitProbes(cfg)
	//
	//if probErr != nil {
	//	logrus.Fatalf("[probes] service started failed %s", probErr.Error())
	//}

	r := service.NewRoom()
	http.Handle("/room", r)
	go r.Run()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), nil); err != nil {
		logrus.Fatalf("Failed to start chat, cause: %s", err.Error())
	}
}
