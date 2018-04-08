package pointfeed

import (
	log "github.com/sirupsen/logrus"
	"sync"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

const bufferSize = 380800 // Size of one full 3d scan at highest resolution.

// PointFeed does pub/sub for points.
type PointFeed struct {
	points sync.Map
	subs   sync.Map
	input  chan coordinates.Point
}

// New makes a new PointFeed and stars it listening for input.
func New() *PointFeed {
	d := &PointFeed{
		points: sync.Map{},
		subs:   sync.Map{},
		input:  make(chan coordinates.Point, bufferSize),
	}
	go d.handleInput()
	return d
}

// Subscribe to the feed. id must be unique.
func (d *PointFeed) Subscribe(id string, c chan coordinates.Point) {
	log.Info(id, " subscribed to point feed.")
	d.subs.Store(id, c)
	go func() {
		d.points.Range(func(key, value interface{}) bool {
			p := key.(coordinates.Point)
			c <- p
			return true
		})
	}()
}

// Unsubscribe from the feed.
func (d *PointFeed) Unsubscribe(id string) {
	d.subs.Delete(id)
}

// Publish a new point.
func (d *PointFeed) Publish(p coordinates.Point) {
	d.input <- p
}

func (d *PointFeed) broadcast(p coordinates.Point) {
	d.subs.Range(func(key, value interface{}) bool {
		c := value.(chan coordinates.Point)
		c <- p
		return true
	})
}

func (d *PointFeed) handleInput() {
	for {
		select {
		case input := <-d.input:
			d.points.Store(input, struct{}{})
			d.broadcast(input)
		}
	}
}
