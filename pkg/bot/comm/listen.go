package comm

import (
	"bufio"
	"encoding/json"
	"github.com/andrewbackes/autonoma/pkg/perception/signal"
	log "github.com/sirupsen/logrus"
)

type SignalHandler func(*signal.Signal)

func (c *Comm) Listen(h SignalHandler) {
	for {
		if c.conn == nil {
			c.connect()
		}
		b, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err != nil {
			log.Error(err)
			c.conn.Close()
			c.conn = nil
		}
		log.Debug("Recieved:", string(b[:len(b)-1]))
		var sig signal.Signal
		err = json.Unmarshal(b, &sig)
		if err == nil {
			h(&sig)
		} else {
			log.Errorf("Could not unmarshal json: %v", err)
		}
	}
}
