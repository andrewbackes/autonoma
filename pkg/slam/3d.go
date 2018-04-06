package slam

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

// ThreeD creates a 3d map.
func ThreeD(g *grid.Grid, bot Bot) {
	log.Info("Mapping...")
	f, _ := os.Create(fmt.Sprintf("output/3d-scan-%v.json", time.Now()))
	for i := -35; i <= 83; i++ {
		scans := bot.LidarScan(i)
		for _, scan := range scans {
			b, err := json.Marshal(scan)
			if err != nil {
				panic(err)
			}
			f.Write(append(b, '\n'))
		}
	}
	f.Close()
	log.Info("Done mapping.")
}
