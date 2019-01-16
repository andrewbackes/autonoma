package signal

type LidarScan struct {
	Orientation Orientation `json:"orientation"`
	Odometer    float64     `json:"odometer"`
	Origin      Vector      `json:"origin"`
	Vectors     []Vector    `json:"vectors"`
}

type Orientation struct {
	Yaw   float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll  float64 `json:"roll"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
