package main

import (
	"fmt"
	"net"
	"os"
)

// Receiver takes information sent from the bot over udp.
type Receiver struct{}

// NewReceiver turns a new Receiver.
func NewReceiver() *Receiver {
	return &Receiver{}
}

// Start begins listening on UDP for data from the bot.
func (r *Receiver) Start() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":9090")
	check(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	check(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// Stop shuts down the receiver.
func (r *Receiver) Stop() {
	// TODO
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
