package main

import (
	"fmt"
	"time"
)

func chooseDir(s *Sensors) string {
	// forward:
	f := !s.IR.Left && !s.IR.Right && s.Echo.Front > 15.0
	if f {
		return "forward"
	}
	ccw := !s.IR.Left && s.IR.Right
	if ccw {
		return "clockwise"
	}
	cw := s.IR.Left && !s.IR.Right
	if cw {
		// TODO
	}
	return "counter_clockwise"
}

func avoidObstacles(b *Bot, s *Sensors) {
	fmt.Println("Avoiding obstacles...")
	for {
		dir := chooseDir(s)
		b.Move(dir, 100*time.Millisecond, 75)
		time.Sleep(100 * time.Millisecond)
	}
}
