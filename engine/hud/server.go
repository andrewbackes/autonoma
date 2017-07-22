package hud

import (
	"github.com/gorilla/mux"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func Serve() {
	m := mux.NewRouter()
	m.HandleFunc("/health", healthCheck).Methods("GET")
	http.ListenAndServe(":80", m)
}
