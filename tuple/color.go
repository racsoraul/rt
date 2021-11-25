package tuple

// NewColor Creates a new color.
func NewColor(r, g, b float64) Tuple {
	return Tuple{r, g, b, 0}
}

// HadamardProduct Blends two colors.
func HadamardProduct(a, b Tuple) Tuple {
	return NewColor(a[0]*b[0], a[1]*b[1], a[2]*b[2])
}
