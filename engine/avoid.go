package main

import (
	"fmt"
	"strings"
	"time"
)

func pattern(h, p []BotMove) bool {
	if len(h) < len(p) {
		return false
	}
	for i := range p {
		if p[len(p)-i-1] != h[len(h)-i-1] {
			return false
		}
	}
	return true
}

func stuck(m []BotMove) bool {
	lrl := []BotMove{
		{dir: "counter_clockwise"},
		{dir: "clockwise"},
		{dir: "counter_clockwise"},
	}
	rlr := []BotMove{
		{dir: "counter_clockwise"},
		{dir: "clockwise"},
		{dir: "counter_clockwise"},
	}
	if pattern(m, lrl) || pattern(m, rlr) {
		return true
	}
	return false
}

func hasntProgressed(m []BotMove) bool {
	c := 10
	if len(m) < c {
		return false
	}
	for i := 0; i < 10; i++ {
		v := m[len(m)-1-i]
		if v.dir == "forward" {
			return false
		}
	}
	return true
}

func bailing(m []BotMove) bool {
	c := 10
	if len(m) < c {
		return false
	}
	b := 0
	for i := 0; i < c; i++ {
		v := m[len(m)-1-i]
		if v.dir == "backward" {
			b++
		}
	}
	if b > 0 && c/b > 2 {
		return true
	}
	return false
}

func chooseDir(s *Sensors, m []BotMove) string {
	if bailing(m) {
		return "counter_clockwise"
	}
	b := s.Echo.Back > 15.0 || s.Echo.Back == None
	if hasntProgressed(m) && b {
		return "backward"
	}
	if stuck(m) {
		return "counter_clockwise"
	}
	f := !s.IR.Left && !s.IR.Right && (s.Echo.Front > 15.0 || s.Echo.Front == None)
	fmt.Println(!s.IR.Left, !s.IR.Right, (s.Echo.Front > 15.0), (s.Echo.Front == None))
	if f {
		return "forward"
	}
	ccw := !s.IR.Left && s.IR.Right
	if ccw {
		return "counter_clockwise"
	}
	cw := s.IR.Left && !s.IR.Right
	if cw {
		return "clockwise"
	}
	if b {
		return "backward"
	}
	return "clockwise"
}

func avoidObstacles(b *Bot, s *Sensors) {
	fmt.Println("Avoiding obstacles...")
	for {
		dir := chooseDir(s, b.moveHistory)
		p := int64(55)
		if strings.Contains(dir, "clockwise") {
			p = 70
		}
		m := BotMove{dir, 100 * time.Millisecond, p}
		fmt.Println(m)
		b.Make(m)
		time.Sleep(150 * time.Millisecond)
	}
}
