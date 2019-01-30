package web

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/pointfeed/subscribers"
)

const chanBufferSize = 380800

type client struct {
	id     string
	done   chan struct{}
	conn   *websocket.Conn
	buffer chan coordinates.Point
}

func newClient(conn *websocket.Conn) *client {
	c := &client{
		id:     strconv.FormatFloat(rand.Float64(), 'E', -1, 64),
		done:   make(chan struct{}),
		conn:   conn,
		buffer: make(chan coordinates.Point, chanBufferSize),
	}
	conn.SetCloseHandler(func(code int, text string) error {
		close(c.done)
		return nil
	})
	log.Info("New websocket client ", *c)
	return c
}

func (c *client) subscribe(sub subscribers.SubscribeUnsubscriber) {
	sub.Subscribe(c.id, c.buffer)
	for {
		select {
		case p := <-c.buffer:
			j, _ := json.Marshal(p)
			c.conn.WriteMessage(websocket.TextMessage, j)
		case <-c.done:
			sub.Unsubscribe(c.id)
			return
		}
	}
}
