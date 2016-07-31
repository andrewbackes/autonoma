package main

import (
	"fmt"
	"strconv"
	"time"
)

const None = -1

type Sensor interface {
	Set(string, string)
}

type IRSensors struct {
	Left  bool
	Right bool
}

func (s IRSensors) String() string {
	return fmt.Sprintf("IR: [left: %t, right: %t]", s.Left, s.Right)
}

func (s *IRSensors) Set(pos, value string) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Print("Couldn't convert to bool:", value)
		return
	}
	switch pos {
	case "left":
		s.Left = b
	case "right":
		s.Right = b
	}
}

type EchoSensors struct {
	Front float64
	Back  float64
}

func (s EchoSensors) String() string {
	trunc := func(f float64) string {
		if f == None {
			return "None"
		}
		return fmt.Sprintf("%.2f", f)
	}
	return fmt.Sprintf("Echo: [front: %s, back: %s]", trunc(s.Front), trunc(s.Back))
}

func (s *EchoSensors) Set(pos, value string) {
	var f float64
	if value == "None" {
		f = None
	} else {
		var err error
		f, err = strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Print("Couldn't convert to float:", value)
			return
		}
	}
	switch pos {
	case "front":
		s.Front = f
	case "back":
		s.Back = f
	}
}

type Compass struct {
	Heading int64
}

func (s Compass) String() string {
	return fmt.Sprintf("Compass: [heading: %d°]", s.Heading)
}

func (s *Compass) Set(pos, value string) {
	switch pos {
	case "heading":
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Print("Couldn't convert to int:", value)
			return
		}
		s.Heading = i
	}
}

type Sensors struct {
	IR      IRSensors
	Echo    EchoSensors
	Compass Compass
	ts      time.Time
}

func (s Sensors) String() string {
	return fmt.Sprint(s.IR, s.Echo, s.Compass)
}

func (s *Sensors) Set(t, pos, value string) {
	sensors := map[string]Sensor{
		"ir":      &s.IR,
		"echo":    &s.Echo,
		"compass": &s.Compass,
	}
	sensors[t].Set(pos, value)
	s.ts = time.Now()
}

func NewSensors() *Sensors {
	return &Sensors{}
}
