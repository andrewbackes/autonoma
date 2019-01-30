package comm

import (
	"encoding/json"
	"github.com/andrewbackes/autonoma/pkg/control"
	log "github.com/sirupsen/logrus"
)

func (c *Comm) Send(cmds control.Commands) {
	if c.conn == nil {
		c.connect()
	}
	err := json.NewEncoder(c.conn).Encode(cmds)
	if err != nil {
		panic(err)
	}
	_, err = c.conn.Write([]byte("\n"))
	if err != nil {
		panic(err)
	}
	log.Debug("Sent:", cmds)
}
