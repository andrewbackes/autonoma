package coordinates

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

func (s CartesianSet) Range(f func(coor Cartesian) bool) {
	for k := range s {
		ret := f(k)
		if ret {
			return
		}
	}
}

func (s CartesianSet) Len() int {
	return len(s)
}
