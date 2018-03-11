package coordinates

import (
	"fmt"
)

type Cartesian struct {
	X, Y int
}

func (c Cartesian) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

type CartesianSet map[Cartesian]struct{}

func NewCartesianSet() CartesianSet {
	return make(map[Cartesian]struct{})
}

// Add a key to a set.
func (s CartesianSet) Add(key Cartesian) {
	s[key] = struct{}{}
}

// Remove a key from a set.
func (s CartesianSet) Remove(key Cartesian) {
	delete(s, key)
}

// Contains looks for keys in a set.
func (s CartesianSet) Contains(keys ...Cartesian) bool {
	for _, key := range keys {
		if _, contains := s[key]; contains {
			return true
		}
	}
	return false
}

// Update combines two sets.
func (s CartesianSet) Update(b CartesianSet) {
	for key := range b {
		s[key] = struct{}{}
	}
}
