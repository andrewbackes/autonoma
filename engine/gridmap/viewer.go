package gridmap

// Viewer can display a map.
type Viewer interface {
	Center() (x, y int)
}
