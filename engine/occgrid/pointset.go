package occgrid

type pointset map[point]struct{}

func (ps pointset) add(p point) {
	ps[p] = struct{}{}
}

func (ps pointset) remove(p point) {
	delete(ps, p)
}

func (ps pointset) contains(p point) bool {
	_, exists := ps[p]
	return exists
}

func newPointset() pointset {
	return map[point]struct{}{}
}
