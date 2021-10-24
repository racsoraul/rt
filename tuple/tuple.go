package tuple

import "fmt"

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
