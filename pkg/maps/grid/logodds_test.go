package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Make sure probability is being converted to and from correctly.
func TestLogOddsProbability(t *testing.T) {
	l := NewLogOdds()
	t.Logf("\nLogOdds: %f\nProbability: %f\n", l, l.Probability())
	assert.InEpsilon(t, initProbability, NewLogOdds().Probability(), 0.01)
}
