package coordinates

type Point struct {
	Origin      Vector `json:"origin"`
	Orientation Euler  `json:"orientation"`
	Vector      Vector `json:"vector"`
}
