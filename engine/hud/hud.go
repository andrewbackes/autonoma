// Package hud is for displaying real time bot data.
package hud

import (
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/gorilla/mux"
	"image/jpeg"
	"image/png"
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

// Serve broadcasts real time data via http.
func (h *Hud) Serve() {
	log.Println("Starting HUD.")
	m := mux.NewRouter()
	m.HandleFunc("/health", healthCheck).Methods("GET")
	m.HandleFunc("/hud.html", showHud).Methods("GET")
	m.HandleFunc("/map.jpeg", h.mapJpegHandler).Methods("GET")
	m.HandleFunc("/map.png", h.mapPngHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", m))
	log.Println("Stopped HUD.")
}

func showHud(w http.ResponseWriter, r *http.Request) {
	log.Println("Showing hud.")
	w.Write([]byte(hudTmpl))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func (h *Hud) mapJpegHandler(w http.ResponseWriter, r *http.Request) {
	err := jpeg.Encode(w, h.mapReader, nil)
	if err != nil {
		panic("Could encode jpeg")
	}
}

func (h *Hud) mapPngHandler(w http.ResponseWriter, r *http.Request) {
	err := png.Encode(w, h.mapReader)
	if err != nil {
		panic("Could encode jpeg")
	}
}

var hudTmpl = `<html>
    <head>
        <style>
            html, body {
                margin: 0;
                padding: 0;
            }
        </style>
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    </head>
    <body>
		<script>
			let mapPath = "/map.png?" + new Date().getTime();
            let init = '<img src="' + mapPath + '">';
            $("body").append(init);
            function refreshImg() {
				let mapPath = "/map.png?" + new Date().getTime();
                let tmpl = '<img style="display: none;" src="' + mapPath + '">';
                $("body").append(tmpl)
                window.setTimeout(function() {
                    $("body").find('img').first().remove();
					$("body").find('img').first().css('display', 'block');
                }, 1000);
            }
            window.setInterval("refreshImg();", 1000);
            
        </script>
    </body>
</html>`
