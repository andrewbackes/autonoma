package main

import (
	"fmt"
	"time"
)

func main() {
	s := NewSensors()
	r := NewReceiver(s)
	go r.Start()
	for {
		fmt.Println(*s)
		time.Sleep(1 * time.Second)
	}
}
