package occgrid

import (
	"github.com/stretchr/testify/assert"
	"image/jpeg"
	"os"
	"testing"
)

func TestIndexConversion(t *testing.T) {
	var assert = assert.New(t)
	g := NewGrid(2, 2)
	assert.Equal(3, g.index(1, 1))
	assert.Equal(2, g.index(0, 1))
	assert.Equal(1, g.index(1, 0))
	assert.Equal(0, g.index(0, 0))

	g10 := NewGrid(10, 10)
	assert.Equal(99, g10.index(9, 9))
}

func TestExampleImage(t *testing.T) {
	g := NewGrid(1000, 1000)
	f, _ := os.Create("occgrid.jpeg")
	err := jpeg.Encode(f, g, nil)
	if err != nil {
		panic("Could encode jpeg")
	}
}
