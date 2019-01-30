package grid

type Odds interface {
	Probability() float64
	Adjust(pmz float64)
}
