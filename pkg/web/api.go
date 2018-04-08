package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/pointfeed/subscribers"
)

// API to serve the front-end.
type API struct {
	data subscribers.SubscribeUnsubscriber
}

// NewAPI server.
func NewAPI(data subscribers.SubscribeUnsubscriber) *API {
	return &API{data: data}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Start serving the API.
func (a *API) Start() {
	log.Info("Serving hud.")
	r := mux.NewRouter()

	r.HandleFunc("/live", a.live)
	r.HandleFunc("/scans/{id}", a.scanByID)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}

func (a *API) live(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client subscribed")
	c := newClient(conn)
	c.subscribe(a.data)
	conn.Close()
	fmt.Println("Client unsubscribed")
}

func (a *API) scanByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client subscribed")
	f, err := os.Open("output/" + id)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	vectors := []coordinates.Vector{}
	for scanner.Scan() {
		var p coordinates.Point
		json.Unmarshal([]byte(scanner.Text()), &p)
		v := coordinates.Vector{X: -p.Vector.X, Y: p.Vector.Z, Z: p.Vector.Y}
		vectors = append(vectors, v)
	}
	for _, v := range vectors {
		j, _ := json.Marshal(v)
		conn.WriteMessage(websocket.TextMessage, j)
		time.Sleep(1 * time.Millisecond)
	}
	conn.Close()
	fmt.Println("Client unsubscribed")
}

/*

r.HandleFunc("/map.png", func(w http.ResponseWriter, r *http.Request) {
		img := (*grid.Image)(g)
		err := png.Encode(w, img)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(err)))
		}
	})
	r.HandleFunc("/vectors", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("output/3d-scan-2018-04-06 18:16:53.807111315 -0700 PDT m=+0.000496495.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		vectors := []coordinates.Vector{}
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			var p coordinates.Point
			json.Unmarshal([]byte(scanner.Text()), &p)
			vectors = append(vectors, p.Vector)
		}
		fmt.Println("Array size:", len(vectors))
		set := map[coordinates.Vector]struct{}{}
		for _, v := range vectors {
			set[v] = struct{}{}
		}
		noDups := []coordinates.Vector{}
		for k := range set {
			noDups = append(noDups, coordinates.Vector{X: -k.X, Y: k.Z, Z: k.Y})
		}
		fmt.Println("No dupes:", len(noDups))
		b, err := json.Marshal(noDups)
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
		// w.Write([]byte(`[{"X": 0, "Y": 0, "Z": 0}]`))
	})
	r.HandleFunc("/2d", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("output/3d-scan-2018-04-06 18:16:53.807111315 -0700 PDT m=+0.000496495.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		vectors := []coordinates.Vector{}
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			var p coordinates.Point
			json.Unmarshal([]byte(scanner.Text()), &p)
			vectors = append(vectors, p.Vector)
		}
		fmt.Println("Array size:", len(vectors))
		set := map[coordinates.Vector]struct{}{}
		for _, v := range vectors {
			set[v] = struct{}{}
		}
		noDups := []coordinates.Vector{}
		for k := range set {
			v := coordinates.Vector{X: -k.X, Y: k.Z, Z: k.Y}
			if v.Y < 20 && v.Y > -10 {
				noDups = append(noDups, v)
			}
		}
		fmt.Println("No dupes:", len(noDups))
		b, err := json.Marshal(noDups)
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	})
*/
