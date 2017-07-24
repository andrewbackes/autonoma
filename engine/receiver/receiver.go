package receiver

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"log"
	"net"
	"os"
	"strings"
)

type Receiver struct {
	mapWriter gridmap.Writer
	sensors   map[string]*sensor.Sensor
}

func New(m gridmap.Writer) *Receiver {
	return &Receiver{
		mapWriter: m,
		sensors:   make(map[string]*sensor.Sensor),
	}
}

func (r *Receiver) Start() {
	fmt.Println("Starting Receiver.")

	//go r.readQueue()

	ServerAddr, err := net.ResolveUDPAddr("udp", ":9090")
	check(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	check(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		//fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		msg := buf[0:n]
		go r.process(msg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	fmt.Println("Stopped Receiver.")
}

func (r *Receiver) process(msg []byte) {
	// TODO: This method kind of sucks.
	s := string(msg)
	if strings.Contains(s, "reading") {
		reading := sensor.DecodeReading(msg)
		occupied, vacant := sensor.Process(r.sensors[reading.SensorID], reading)
		for o := range occupied {
			r.mapWriter.Occupied(o.X, o.Y)
		}
		for v := range vacant {
			r.mapWriter.Vacant(v.X, v.Y)
		}
	} else if strings.Contains(s, "sensor") {
		s := sensor.DecodeSensor(msg)
		log.Println("Registering sensor:", s)
		r.sensors[s.ID] = s
	}
}

/*
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
*/
func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

/*
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


// Stop shuts down the receiver.
func (r *Receiver) Stop() {
	// TODO
}

*/
