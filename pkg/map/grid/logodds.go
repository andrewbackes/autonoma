package grid

import (
	"math"
)

/*
	- The proposition occ(i,j) means:
		– The cell Cij is occupied.
	- Probability: p(occ(i,j)) has range [0,1].
	- Odds: o(occ(i,j)) has range [0,+∞).
		- o(A) = p(A)/p(¬A)
	- Log odds: log o(occ(i,j)) has range (−∞,+∞)
	- Each cell Cij holds the value log o(occ(i,j))
		– Cij = 0 corresponds to p(occ(i,j)) = 0.5

*/

type LogOdds float64

var (
	initProbability = 0.2 // 0.2 to 0.5 depending on expected obstacle density
	initLogOdds     = LogOdds(math.Log(initProbability) - math.Log(1-initProbability))
)

func NewLogOdds() Odds {
	return initLogOdds
}
func (l LogOdds) Probability() float64 {
	p := 1 - (1 / math.Exp(float64(l)))
	if p > 1 {
		return 1
	}
	if p < 0 {
		return 0
	}
	return p
}

// Adjust by probability of M given Z [ Or p(m|z) ].
func (l LogOdds) Adjust(pmz float64) {
	// l = l + log p(m|z) - log (1 - p(m|z)) - log p(m) + log(1 - p(m))
	pm := l.Probability()
	lp := float64(l) + math.Log(pmz) - math.Log(1-pmz) - math.Log(pm) + math.Log(1-pm)
	l = LogOdds(lp)
}
