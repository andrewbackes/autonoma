package main

import (
	"fmt"
	"time"
)

func main() {
	s := NewSensors()
	r := NewReceiver(s)
	go r.Start()
	go func() {
		for {
			fmt.Println(*s)
			time.Sleep(1 * time.Second)
		}
	}()
	b := NewBot()
	avoidObstacles(b, s)
}
