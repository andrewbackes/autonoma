package coordinates

import (
	"sync"
)

type CartesianSet struct {
	set   map[Cartesian]struct{}
	mutex *sync.RWMutex
}

func NewCartesianSet() CartesianSet {
	return CartesianSet{
		set:   make(map[Cartesian]struct{}),
		mutex: &sync.RWMutex{},
	}
}

// Add a key to a set.
func (s CartesianSet) Add(key Cartesian) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.set[key] = struct{}{}

}

// Remove a key from a set.
func (s CartesianSet) Remove(key Cartesian) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.set, key)
}

// Contains looks for keys in a set.
func (s CartesianSet) Contains(keys ...Cartesian) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, key := range keys {
		if _, contains := s.set[key]; contains {
			return true
		}
	}
	return false
}

func (s CartesianSet) Range(f func(coor Cartesian) bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for k := range s.set {
		ret := f(k)
		if ret {
			return
		}
	}
}

func (s CartesianSet) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.set)
}
