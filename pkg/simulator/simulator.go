// Package simulator is for simulating communication with a bot.
package simulator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/control"
	"github.com/andrewbackes/autonoma/pkg/perception/signal"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

// Simulator mimics a robot.
type Simulator struct {
	conn          net.Conn
	done          bool
	sequence      []signal.Signal
	sequenceDelay time.Duration
}

// New Simulator.
func New(sequenceFile string, sequenceDelay time.Duration) *Simulator {
	s := &Simulator{
		done:          false,
		sequenceDelay: sequenceDelay,
	}
	f, err := os.Open(sequenceFile)
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&s.sequence)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Simulator) sendSequence() {
	fmt.Println("Sending sequence.")
	for _, sig := range s.sequence {
		b, err := json.Marshal(sig)
		if err != nil {
			panic(err)
		}
		b = append(b, '\n')
		_, err = s.conn.Write(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sent. Waiting...")
		time.Sleep(s.sequenceDelay)
	}
}

// Listen for incoming connections.
func (s *Simulator) Listen() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Listening...")
		s.conn, err = listener.Accept()
		fmt.Println("Accepted connection.")
		if err != nil {
			panic(err)
		}
		go s.sendSequence()
		s.read()
	}
}

func (s *Simulator) read() {
	for {
		b, err := bufio.NewReader(s.conn).ReadBytes('\n')
		if err != nil {
			log.Error(err)
			s.conn.Close()
			s.conn = nil
			return
		}
		log.Debug("Recieved:", string(b[:len(b)-1]))
		var commands control.Commands
		err = json.Unmarshal(b, &commands)
		if err != nil {
			log.Errorf("Could not unmarshal json: %v", err)
		}
	}
}
