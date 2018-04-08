package slam

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

// Manual is used to manually reposition the robot while it takes sensor readings.
func Manual(g *grid.Grid, bot bot) {
	log.Info("Manually Mapping...")
	readings := make([][]sensor.Reading, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Face forward:
		log.Info("Press enter to begin scan. Type 'exit' to finish.")
		scanner.Scan()
		if scanner.Text() == "exit" {
			break
		}
		log.Info("Scanning...")
		set1 := bot.Scan()
		for i := range set1 {
			// scans go from -90 to 90
			set1[i].RelativeHeading += 90
		}
		// Face backward:
		log.Info("Turn the bot around and press 'enter'")
		scanner.Scan()
		log.Info("Scanning...")
		set2 := bot.Scan()
		for i := range set2 {
			// scans go from -90 to 90
			set2[i].RelativeHeading += 270
		}
		log.Info("Disguard? [no]")
		scanner.Scan()
		if scanner.Text() != "yes" {
			readings = append(readings, append(set1, set2...))
		}
	}
	// Save
	log.Info("Finished Scanning.")
	data, err := json.Marshal(readings)
	if err != nil {
		panic(err)
	}
	log.Info("Saving text file.")
	txt, _ := os.Create(fmt.Sprintf("output/scan-%v", time.Now()))
	txt.Write(data)
	txt.Close()
}
