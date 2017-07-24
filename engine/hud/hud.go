package hud

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/gorilla/mux"
	"image/jpeg"
	"log"
	"net/http"
)

type Hud struct {
	mapReader gridmap.Reader
}

func New(r gridmap.Reader) *Hud {
	return &Hud{
		mapReader: r,
	}
}

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
