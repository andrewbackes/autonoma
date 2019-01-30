package set

// Set is a collection of keys.
type Set map[string]struct{}

// New constructs a new set.
func New() Set {
	return make(map[string]struct{})
}

// Add a key to a set.
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

// Remove a key from a set.
func (s Set) Remove(key string) {
	delete(s, key)
}

// Contains looks for keys in a set.
func (s Set) Contains(keys ...string) bool {
	for _, key := range keys {
		if _, contains := s[key]; contains {
			return true
		}
	}
	return false
}

// Update combines two sets.
func (s Set) Update(b Set) {
	for key := range b {
		s[key] = struct{}{}
	}
}
