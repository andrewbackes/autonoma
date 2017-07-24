// Package hud is for displaying real time bot data.
package hud

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/gorilla/mux"
	"image/jpeg"
	"log"
	"net/http"
)

// Hud is a Heads Up Display. It is for showing real time bot data.
type Hud struct {
	mapReader gridmap.Reader
}

// New creates a Hud.
func New(r gridmap.Reader) *Hud {
	return &Hud{
		mapReader: r,
	}
}

// Start broadcasts real time data via http.
func (h *Hud) Start() {
	fmt.Println("Starting HUD.")
	m := mux.NewRouter()
	m.HandleFunc("/health", healthCheck).Methods("GET")
	m.HandleFunc("/map.jpeg", h.mapHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", m))
	fmt.Println("Stopped HUD.")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func (h *Hud) mapHandler(w http.ResponseWriter, r *http.Request) {
	err := jpeg.Encode(w, h.mapReader, nil)
	if err != nil {
		panic("Could encode jpeg")
	}
}
