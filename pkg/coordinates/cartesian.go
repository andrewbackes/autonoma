package coordinates

import (
	"fmt"
)

type Cartesian struct {
	X, Y int64
}

func (c Cartesian) String() string {
	return fmt.Sprintf(`%d,%d`, c.X, c.Y)
}
