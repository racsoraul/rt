package tuple

import (
	"math"
	"rt"
)

// NewPoint Creates a new Point.
func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// IsPoint Checks if tuple is a Point.
func (t Tuple) IsPoint() bool {
	return math.Abs(t.W-1.0) < rt.EPSILON
}
