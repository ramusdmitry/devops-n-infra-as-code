package probes

import (
	chatApp "chat-app-service"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Probes struct {
	isAlive bool
	isReady bool
}

func NewProbes() *Probes {
	return &Probes{
		isReady: false,
		isAlive: false,
	}
}

func (p *Probes) InitProbes(config *chatApp.Config) error {

	var err, probErr error

	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"status": "ok"}`))
		if err != nil {
			logrus.Errorf("failed to write response: %s", err.Error())
		}
		p.isAlive = true
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logrus.Errorf("failed to connect tiny websockets server%s", err.Error())
			p.isReady = false
		}
		defer conn.Close()
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		dialer := websocket.DefaultDialer
		dialer.HandshakeTimeout = time.Second * 5
		conn, _, err := dialer.Dial("ws://localhost:8042/", nil)
		if err != nil {
			logrus.Errorf("failed check connection to websockets with probe: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			p.isReady = false
		}
		defer conn.Close()

		err = conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			logrus.Warnf("failed send message to websockets with probe: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			p.isReady = false
		}

		w.WriteHeader(http.StatusOK)
		p.isReady = true
	})

	go func() {

		err := http.ListenAndServe(":8042", nil)
		if err != nil {
			logrus.Errorf("start probes failed: %s", err.Error())
		}

		if !p.isAlive {
			logrus.Errorf("connection unavailable")
			probErr = errors.New("status check failed")
		}
		if !p.isReady {
			logrus.Errorf("websocket server unavailable")
			probErr = errors.New("status check failed")
		}
	}()

	return probErr
}
