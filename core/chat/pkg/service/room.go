package service

import (
	chatApp "chat-app-service/pkg/model"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 1024
	limitMessages     = 2
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: messageBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	forward   chan *chatApp.Message
	join      chan *client
	leave     chan *client
	clients   map[*client]bool
	history   []chatApp.Message
	timeStart int64
	timeLimit int64
}

func (r *Room) Run() {

	for {
		select {

		case client := <-r.join:
			logrus.Info("[Chat] a new connection")
			r.clients[client] = true

		case client := <-r.leave:
			logrus.Info("[Chat] user left chat")
			delete(r.clients, client)
			close(client.send)

		case msg := <-r.forward:
			currentTime := time.Now().Unix()
			for client := range r.clients {
				select {
				case client.send <- msg:
					if client.currentCountMessages < limitMessages {
						r.history = append(r.history, *msg)
						client.currentCountMessages += 1
					} else {
						if currentTime > r.timeStart+r.timeLimit {
							r.timeStart = currentTime
							client.currentCountMessages = 0
							r.history = append(r.history, *msg)
						} else {
							go func() {
								client.mu.Lock()
								err := client.socket.WriteJSON("channel is filled")
								if err != nil {
									logrus.Errorf("[Chat] error when channel was filled: %s", err.Error())
									return
								}
								defer client.mu.Unlock()
							}()
						}
					}
					LoadMessages(r)
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}

		}
	}

}

func LoadMessages(r *Room) {
	for _, i := range r.history {
		logrus.Printf("[%s] %s: %s %d", i.When.Format("2006-01-02 15:04:05.000000"), i.Name, i.Message, len(r.history))
	}
}

func NewRoom() *Room {
	return &Room{
		forward:   make(chan *chatApp.Message),
		join:      make(chan *client),
		leave:     make(chan *client),
		clients:   make(map[*client]bool),
		history:   make([]chatApp.Message, 0),
		timeStart: time.Now().Unix(),
		timeLimit: 30, // 24h = 86400s
	}
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.Fatalf("[Chat] ServeHTTP crashed, cause %s", err.Error())
		return
	}
	client := &client{
		socket:               socket,
		send:                 make(chan *chatApp.Message, messageBufferSize),
		currentCountMessages: 0,
		room:                 r,
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.Write()
	client.Read()
}
