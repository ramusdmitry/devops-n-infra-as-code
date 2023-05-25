package service

import (
	chatApp "chat-app-service/pkg/model"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type client struct {
	socket               *websocket.Conn
	send                 chan *chatApp.Message
	currentCountMessages int
	room                 *Room
	mu                   sync.Mutex
}

func (c *client) Read() {
	for {
		var msg *chatApp.Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			c.room.forward <- msg
		} else {
			logrus.Warnf("[Chat] failed to parse JSON of message, cause: %s", err.Error())
			break
		}
	}
}

func (c *client) Write() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			logrus.Warnf("[Chat] failed to write JSON of message, cause: %s", err.Error())
			break
		}
	}
}

func (c *client) WriteMessage(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.socket.WriteJSON(message); err != nil {
		logrus.Warnf("[Chat] failed to write JSON of message, cause: %s", err.Error())
		return
	}
}
