package receiver

import (
	"bufio"
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
	if strings.Contains(msg, "sensorId") {
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
	} else if strings.Contains(msg, "\"id\"") {
		s := sensor.DecodeSensor([]byte(msg))
		log.Println("Registering sensor:", s)
		r.sensors[s.ID] = s
	}
}

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
}
