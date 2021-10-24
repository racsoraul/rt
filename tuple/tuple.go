package tuple

import (
	"fmt"
	"math"
	"rt"
)

// Tuple underlying structure for vector and points.
type Tuple struct {
	X, Y, Z, W float64
}

// NewTuple Creates a new Tuple.
func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func (t Tuple) String() string {
	return fmt.Sprintf("X: %g, Y: %g, Z: %g, W: %g", t.X, t.Y, t.Z, t.W)
}

func (t Tuple) IsEqual(to Tuple) bool {
	return equalFloat(t.X, to.X) && equalFloat(t.Y, to.Y) && equalFloat(t.Z, to.Z) && equalFloat(t.W, to.W)
}

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < rt.EPSILON
}
