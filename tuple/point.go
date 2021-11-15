package tuple

// NewPoint Creates a new Point.
func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// IsPoint Checks if tuple is a Point.
func (t Tuple) IsPoint() bool {
	return equalFloat(t[3], 1)
}
