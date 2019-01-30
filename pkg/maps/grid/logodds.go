package grid

import (
	// log "github.com/sirupsen/logrus"

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

type LogOdds struct {
	odds float64
}

var (
	initProbability = 0.2 // 0.2 to 0.5 depending on expected obstacle density
	initLogOdds     = math.Log(initProbability) - math.Log(1-initProbability)
	maxLogOdds      = math.Log(0.99) - math.Log(1-0.99)
	minLogOdds      = math.Log(0.01) - math.Log(1-0.01)
)

func NewLogOdds() Odds {
	return &LogOdds{
		odds: initLogOdds,
	}
}

// Probability has range [0,1].
func (l *LogOdds) Probability() float64 {
	p := 1 - (1 / (1 + math.Exp(l.odds)))
	return p
}

// Adjust by probability of M given Z [ Or p(m|z) ].
func (l *LogOdds) Adjust(pmz float64) {
	// l = l + log p(m|z) - log (1 - p(m|z)) - log p(m) + log(1 - p(m))
	pm := l.Probability()
	lp := l.odds + math.Log(pmz) - math.Log(1-pmz) - math.Log(pm) + math.Log(1-pm)
	l.odds = lp
	// log.Infof("Updating logodds %f to %f probability %f to %f", l.odds, lp, pm, l.Probability())
}
