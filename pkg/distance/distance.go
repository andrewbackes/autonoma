package distance

type Distance float64

const (
	Centimeter Distance = 1
	Meter      Distance = 1000
)

func (d Distance) Floor(unit Distance) Distance {
	return Distance(int(d / unit))
}
