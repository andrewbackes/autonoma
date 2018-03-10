package grid

import (
	"testing"
)

func TestLogOddsAdjust(t *testing.T) {
	l := NewLogOdds()
	t.Logf("\n%f %f\n", l, l.Probability())
	t.Fail()
}
