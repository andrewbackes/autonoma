package sensor

type Location struct {
	X, Y int
}

type LocationSet map[Location]struct{}

func (s LocationSet) Add(l Location) {
	s[l] = struct{}{}
}

func (s LocationSet) Remove(l Location) {
	delete(s, l)
}

func (s LocationSet) Contains(l Location) bool {
	_, exists := s[l]
	return exists
}

func NewLocationSet() LocationSet {
	return map[Location]struct{}{}
}
