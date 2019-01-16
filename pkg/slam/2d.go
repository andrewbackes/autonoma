package slam

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// TwoD creates a 3d map.
func TwoD(bot Bot) {
	log.Info("Mapping...")
	f, _ := os.Create(fmt.Sprintf("output/2d-scan-%v.json", time.Now()))
	scans := bot.LidarScan(0)
	for _, scan := range scans {
		b, err := json.Marshal(scan)
		if err != nil {
			panic(err)
		}
		f.Write(append(b, '\n'))
	}

	f.Close()
	log.Info("Done mapping.")
}
