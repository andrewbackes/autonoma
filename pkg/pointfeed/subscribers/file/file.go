// Package file subscribes to a pointfeed and saves its contents to a file.
package file

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"os"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/pointfeed/subscribers"
)

const bufferSize = 380800

type file struct {
	filename string
	f        io.Writer
	buffer   chan coordinates.Point
	done     chan struct{}
}

func newFile(filename string) (*file, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &file{
		filename: filename,
		f:        f,
		buffer:   make(chan coordinates.Point, bufferSize),
		done:     make(chan struct{}),
	}, nil
}

func (f *file) subscribe(to subscribers.SubscribeUnsubscriber) {
	to.Subscribe(f.filename, f.buffer)
	for {
		select {
		case p := <-f.buffer:
			f.write(p)
		case <-f.done:
			to.Unsubscribe(f.filename)
			return
		}
	}
}

func (f *file) write(p coordinates.Point) {
	j, err := json.Marshal(p)
	if err != nil {
		log.Error("Could not write to file - ", err)
		return
	}
	f.f.Write(append(j, '\n'))
}

func (f *file) Close() {
	close(f.done)
}

// Subscribe to a feed and write the feed to a file.
func Subscribe(filename string, to subscribers.SubscribeUnsubscriber) {
	f, err := newFile(filename)
	if err != nil {
		log.Error("Could not subscribe file to point feed - ", err)
		return
	}
	defer f.Close()
	f.subscribe(to)
}
