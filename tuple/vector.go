package tuple

// NewVector Creates a new Vector.
func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// IsVector Checks if tuple is a Vector.
func (t Tuple) IsVector() bool {
	return EqualFloat(t[3], 0)
}
