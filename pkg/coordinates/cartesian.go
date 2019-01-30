package coordinates

import (
	"fmt"
)

type Cartesian struct {
	X, Y int
}

func (c Cartesian) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func Add(c Cartesian, v interface{}) Cartesian {
	if cart, isCart := v.(Cartesian); isCart {
		return Cartesian{
			X: c.X + cart.X,
			Y: c.Y + cart.Y,
		}
	} else if comp, isComp := v.(CompassRose); isComp {
		asCart := comp.Cartesian()
		return Cartesian{
			X: c.X + asCart.X,
			Y: c.Y + asCart.Y,
		}
	}
	panic("Adding vectors requires either compass rose or cartesian coordinates.")
}
