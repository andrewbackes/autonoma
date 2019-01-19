// Package ui declares the user interface. In this case it is an HTTP API.
package ui

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/andrewbackes/autonoma/pkg/perception"
)

type UI struct {
	p perceiver
}

func New(p perceiver) *UI {
	return &UI{p: p}
}

func (ui *UI) ListenAndServe() {
	log.Info("Starting web ui.")
	http.HandleFunc("/perception", ui.perceptionHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Could not start web UI - ", err)
	}
}

type perceiver interface {
	Perception() *perception.Perception
}

func (ui *UI) perceptionHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(ui.p.Perception())
	if err != nil {
		fmt.Printf("error - %v", err)
	}
}
