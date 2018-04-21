package coordinates

import (
	"sync"
)

type VectorSet struct {
	set   map[Vector]struct{}
	mutex *sync.RWMutex
}

func NewVectorSet() VectorSet {
	return VectorSet{
		set:   make(map[Vector]struct{}),
		mutex: &sync.RWMutex{},
	}
}

// Add a key to a set.
func (s VectorSet) Add(key Vector) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.set[key] = struct{}{}

}

// Remove a key from a set.
func (s VectorSet) Remove(key Vector) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.set, key)
}

// Contains looks for keys in a set.
func (s VectorSet) Contains(keys ...Vector) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, key := range keys {
		if _, contains := s.set[key]; contains {
			return true
		}
	}
	return false
}

func (s VectorSet) Range(f func(coor Vector) bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for k := range s.set {
		ret := f(k)
		if ret {
			return
		}
	}
}

func (s VectorSet) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.set)
}
