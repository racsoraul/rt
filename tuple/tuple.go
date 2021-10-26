package tuple

import (
	"errors"
	"fmt"
	"math"
	"rt"
)

var (
	ErrorInvalidAddOp = errors.New("invalid Add operation. Point + Point is not allowed")
	ErrorInvalidSubOp = errors.New("invalid Sub operation. Vector - Point is not allowed")
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
	return fmt.Sprintf("[X: %g, Y: %g, Z: %g, W: %g]", t.X, t.Y, t.Z, t.W)
}

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < rt.EPSILON
}

func (t Tuple) IsEqual(to Tuple) bool {
	return equalFloat(t.X, to.X) && equalFloat(t.Y, to.Y) && equalFloat(t.Z, to.Z) && equalFloat(t.W, to.W)
}

// Add Performs addition of two tuples. Valid operations are
// Vector + Point (Point + Vector) and Vector + Vector. When attempting
// to add Point + Point r will be the zero-value of Tuple and
// err will be ErrorInvalidAddOp.
func Add(a, b Tuple) (r Tuple, err error) {
	if a.IsPoint() && b.IsPoint() {
		return Tuple{}, ErrorInvalidAddOp
	}
	return NewTuple(a.X+b.X, a.Y+b.Y, a.Z+b.Z, a.W+b.W), nil
}

// Sub Performs subtraction of two tuples. Valid operations are
// Point - Point, Point - Vector and Vector - Vector. When attempting
// to subtract Vector - Point r will be zero-value of Tuple and
// err will be ErrorInvalidSubOp.
func Sub(a, b Tuple) (r Tuple, err error) {
	if a.IsVector() && b.IsPoint() {
		return Tuple{}, ErrorInvalidSubOp
	}
	return NewTuple(a.X-b.X, a.Y-b.Y, a.Z-b.Z, a.W-b.W), nil
}
