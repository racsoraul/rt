package tuple

import (
	"errors"
	"fmt"
	"math"
	"rt"
)

var (
	ErrorInvalidAddOp   = errors.New("invalid Add operation. Point + Point is not allowed")
	ErrorInvalidSubOp   = errors.New("invalid Sub operation. Vector - Point is not allowed")
	ErrorDivisionByZero = errors.New("division by zero")
)

// Tuple underlying structure for vector and points.
type Tuple [4]float64

// NewTuple Creates a new Tuple.
func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func (t Tuple) String() string {
	return fmt.Sprintf("[X: %g, Y: %g, Z: %g, W: %g]", t[0], t[1], t[2], t[3])
}

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < rt.EPSILON
}

func (t Tuple) IsEqual(to Tuple) bool {
	return equalFloat(t[0], to[0]) && equalFloat(t[1], to[1]) && equalFloat(t[2], to[2]) && equalFloat(t[3], to[3])
}

// Add Performs addition of two tuples. Valid operations are
// Vector + Point (Point + Vector) and Vector + Vector. When attempting
// to add Point + Point r will be the zero-value of Tuple and
// err will be ErrorInvalidAddOp.
func Add(a, b Tuple) (r Tuple, err error) {
	if a.IsPoint() && b.IsPoint() {
		return Tuple{}, ErrorInvalidAddOp
	}
	return NewTuple(a[0]+b[0], a[1]+b[1], a[2]+b[2], a[3]+b[3]), nil
}

// Sub Performs subtraction of two tuples. Valid operations are
// Point - Point, Point - Vector and Vector - Vector. When attempting
// to subtract Vector - Point r will be zero-value of Tuple and
// err will be ErrorInvalidSubOp.
func Sub(a, b Tuple) (r Tuple, err error) {
	if a.IsVector() && b.IsPoint() {
		return Tuple{}, ErrorInvalidSubOp
	}
	return NewTuple(a[0]-b[0], a[1]-b[1], a[2]-b[2], a[3]-b[3]), nil
}

// Neg Negates a tuple.
func Neg(v Tuple) Tuple {
	return NewTuple(-v[0], -v[1], -v[2], -v[3])
}

// Scale Scales a tuple by a scalar value.
func Scale(t Tuple, s float64) Tuple {
	return NewTuple(t[0]*s, t[1]*s, t[2]*s, t[3]*s)
}

// Div Divides a tuple by a scalar value. Returns ErrorDivisionByZero
// when s is 0.
func Div(t Tuple, s float64) (Tuple, error) {
	if s == 0 {
		return Tuple{}, ErrorDivisionByZero
	}
	return NewTuple(t[0]/s, t[1]/s, t[2]/s, t[3]/s), nil
}

// Mag Returns the magnitude of the vector.
func (t Tuple) Mag() float64 {
	if t.IsPoint() {
		return 0
	}
	return math.Sqrt(math.Pow(t[0], 2) + math.Pow(t[1], 2) + math.Pow(t[2], 2))
}

// Normalize Normalizes a vector.
func Normalize(t Tuple) Tuple {
	return NewVector(t[0]/t.Mag(), t[1]/t.Mag(), t[2]/t.Mag())
}

// Dot Performs dot product between two vectors.
func Dot(a, b Tuple) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Cross Performs cross product between two vectors.
func Cross(a, b Tuple) Tuple {
	return NewVector(a[1]*b[2]-a[2]*b[1], a[2]*b[0]-a[0]*b[2], a[0]*b[1]-a[1]*b[0])
}
