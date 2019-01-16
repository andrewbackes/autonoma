// Package ui declares the user interface. In this case it is an HTTP API.
package ui

import (
	"encoding/json"
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/perception"
	"net/http"
)

type UI struct {
	p perceiver
}

func New(p perceiver) *UI {
	return &UI{p: p}
}

func (ui *UI) ListenAndServe() {
	http.HandleFunc("/perception", ui.perceptionHandler)
	http.ListenAndServe(":8080", nil)
}

type perceiver interface {
	Perception() *perception.Perception
}

func ListenAndServe(p perceiver) {

}

func (ui *UI) perceptionHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(ui.p.Perception())
	if err != nil {
		fmt.Printf("error - %v", err)
	}
}
