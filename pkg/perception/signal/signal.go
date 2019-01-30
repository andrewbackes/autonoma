package signal

// Signal is a generic event sent from a sensor reading.
type Signal struct {
	Type  string      `json:"type"`
	Event interface{} `json:"event"`
}
