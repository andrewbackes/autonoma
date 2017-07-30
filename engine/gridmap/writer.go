package gridmap

// Writer is an interface for creating maps.
type Writer interface {
	Occupied(x, y int)
	Vacant(x, y int)
	Path(x, y int)
	SetPosition(x, y int)
}
