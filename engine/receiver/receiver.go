package receiver

import (
	"bufio"
	"encoding/json"
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"log"
	"net"
	"os"
	"strings"
)

// Receiver processes sensor data.
type Receiver struct {
	mapWriter gridmap.Writer
	sensors   map[string]*sensor.Sensor
}

// New creates a Receiver.
func New(m gridmap.Writer) *Receiver {
	return &Receiver{
		mapWriter: m,
		sensors:   make(map[string]*sensor.Sensor),
	}
}

// Listen begins listening for data over UDP.
func (r *Receiver) Listen(conn net.Conn) {
	log.Println("Starting Receiver.")
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		go r.process(strings.TrimRight(msg, "\n"))
	}
	log.Println("Stopped Receiver.")
}

func (r *Receiver) process(msg string) {
	// TODO: This method kind of sucks.
	if msg[:len("READING")] == "READING" {
		msg = msg[len("READING"):]
		log.Println("Received Reading", msg)
		reading := sensor.DecodeReading([]byte(msg))
		if reading.SensorID != "compass" {
			occupied, vacant := sensor.Process(r.sensors[reading.SensorID], reading)
			for o := range occupied {
				r.mapWriter.Occupied(o.X, o.Y)
			}
			for v := range vacant {
				r.mapWriter.Vacant(v.X, v.Y)
			}
		}
	} else if msg[:len("SENSOR")] == "SENSOR" {
		msg = msg[len("SENSOR"):]
		s := sensor.DecodeSensor([]byte(msg))
		log.Println("Registering sensor:", s)
		r.sensors[s.ID] = s
	} else if msg[:len("LOCATION")] == "LOCATION" {
		msg = msg[len("LOCATION"):]
		loc := &sensor.Location{}
		if err := json.Unmarshal([]byte(msg), &loc); err == nil {
			r.mapWriter.Path(loc.X, loc.Y)
		} else {
			log.Println("Could not decode", string(msg), err)
		}
	} else {
		log.Println("Could not categorize:", msg)
	}
}

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
}
