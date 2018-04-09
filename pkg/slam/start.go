package slam

import (
	log "github.com/sirupsen/logrus"
)

// Start SLAM.
func Start(bot Bot) {
	return
	log.Info("Mapping...")
	for i := -35; i <= 83; i++ {
		bot.LidarScan(i)
	}
	log.Info("Done mapping.")
}
