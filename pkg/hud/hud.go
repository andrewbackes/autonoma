package hud

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"image/png"
	"net/http"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

func ListenAndServe(g *grid.Grid) {
	log.Info("Serving hud.")
	http.HandleFunc("/map.png", func(w http.ResponseWriter, r *http.Request) {
		err := png.Encode(w, grid.Image(*g))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(err)))
		}
	})
	http.ListenAndServe(":8080", nil)
}
