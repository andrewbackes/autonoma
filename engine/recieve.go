package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Receiver takes information sent from the bot over udp.
type Receiver struct {
	s *Sensors
	q chan string
}

// NewReceiver turns a new Receiver.
func NewReceiver(s *Sensors) *Receiver {
	return &Receiver{
		s: s,
		q: make(chan string, 1024),
	}
}

// Start begins listening on UDP for data from the bot.
func (r *Receiver) Start() {
	fmt.Println("Starting Receiver.")
	go r.readQueue()

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":9090")
	check(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	check(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		//fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		r.q <- string(buf[0:n])
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// Stop shuts down the receiver.
func (r *Receiver) Stop() {
	// TODO
}

func (r *Receiver) readQueue() {
	for reading := range r.q {
		//fmt.Println(reading)
		spl := strings.Split(reading, " ")
		if len(spl) >= 3 {
			sensor, pos, value := spl[0], spl[1], spl[2]
			r.s.Set(sensor, pos, value)
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
