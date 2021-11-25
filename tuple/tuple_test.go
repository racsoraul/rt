package tuple

import (
	"math"
	"testing"
)

func TestIsPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	if !equalFloat(p[0], 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", p[0])
	}
	if !equalFloat(p[1], -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", p[1])
	}
	if !equalFloat(p[2], 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", p[2])
	}
	if !equalFloat(p[3], 1.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 1\n", p[3])
	}
	if !p.IsPoint() {
		t.Fatalf("tuple is a vector. It was expected to be a point: %v", p)
	}
}

func TestIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	if !equalFloat(v[0], 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", v[0])
	}
	if !equalFloat(v[1], -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", v[1])
	}
	if !equalFloat(v[2], 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", v[2])
	}
	if !equalFloat(v[3], 0.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 0\n", v[3])
	}
	if !v.IsVector() {
		t.Fatalf("tuple is a point. It was expected to be a vector: %v", v)
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name                     string
		inputA, inputB, expected Tuple
		err                      error
	}{
		{
			"Point + Vector = Point",
			NewPoint(3, -2, 5),
			NewVector(-2, 3, 1),
			NewPoint(1, 1, 6),
			nil,
		},
		{
			"Vector + Vector = Vector",
			NewVector(3, -2, 5),
			NewVector(-2, 3, 1),
			NewVector(1, 1, 6),
			nil,
		},
		{
			"Point + Point = Error",
			NewPoint(3, -2, 5),
			NewPoint(6, 9, 1),
			NewTuple(0, 0, 0, 0),
			ErrorInvalidAddOp,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Add(tc.inputA, tc.inputB)
			if err != tc.err {
				t.Fatalf("error adding tuples. Unexpected err value: %v; wanted: %v", err, tc.err)
			}
			if !result.IsEqual(tc.expected) {
				t.Errorf("got: %v; want: %v", result, tc.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	testCases := []struct {
		name                     string
		inputA, inputB, expected Tuple
		err                      error
	}{
		{
			"Point - Point = Vector",
			NewPoint(3, 2, 1),
			NewPoint(5, 6, 7),
			NewVector(-2, -4, -6),
			nil,
		},
		{
			"Point - Vector = Point",
			NewPoint(3, 2, 1),
			NewVector(5, 6, 7),
			NewPoint(-2, -4, -6),
			nil,
		},
		{
			"Vector - Vector = Vector",
			NewVector(3, 2, 1),
			NewVector(5, 6, 7),
			NewVector(-2, -4, -6),
			nil,
		},
		{
			"Vector - Point = Error",
			NewVector(3, 2, 1),
			NewPoint(5, 6, 7),
			NewTuple(0, 0, 0, 0),
			ErrorInvalidSubOp,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Sub(tc.inputA, tc.inputB)
			if err != tc.err {
				t.Fatalf("error subtracting tuples. Unexpected err value: %v; wanted: %v", err, tc.err)
			}
			if !result.IsEqual(tc.expected) {
				t.Errorf("got: %v; want: %v", result, tc.expected)
			}
		})
	}
}

func TestNeg(t *testing.T) {
	v := NewTuple(1, -2, 3, -4)
	expected := NewTuple(-1, 2, -3, 4)
	actual := Neg(v)
	if !actual.IsEqual(expected) {
		t.Fatalf("got: %v; want: %v", actual, expected)
	}
}

func TestScale(t *testing.T) {
	testCases := []struct {
		name     string
		inputA   Tuple
		inputB   float64
		expected Tuple
	}{
		{
			"Multiplying tuple by a scalar",
			NewTuple(1, -2, 3, -4),
			3.5,
			NewTuple(3.5, -7, 10.5, -14),
		},
		{
			"Multiplying tuple by a fraction",
			NewTuple(1, -2, 3, -4),
			0.5,
			NewTuple(0.5, -1, 1.5, -2),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Scale(tc.inputA, tc.inputB)
			if !actual.IsEqual(tc.expected) {
				t.Fatalf("got: %v; want: %v", actual, tc.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)
	actual, err := Div(tuple, 2)
	if err != nil {
		t.Fatalf("Unexpected division by zero")
	}
	expected := NewTuple(0.5, -1, 1.5, -2)
	if !actual.IsEqual(expected) {
		t.Fatalf("got: %v; want: %v", actual, expected)
	}
}

func TestMag(t *testing.T) {
	testCases := []struct {
		name     string
		input    Tuple
		expected float64
	}{
		{
			"magnitude of v(1,0,0)",
			NewVector(1, 0, 0),
			1,
		},
		{
			"magnitude of v(0,1,0)",
			NewVector(0, 1, 0),
			1,
		},
		{
			"magnitude of v(0,0,1)",
			NewVector(0, 0, 1),
			1,
		},
		{
			"magnitude of v(1,2,3)",
			NewVector(1, 2, 3),
			math.Sqrt(14),
		},
		{
			"magnitude of v(-1,-2,-3)",
			NewVector(-1, -2, -3),
			math.Sqrt(14),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.Mag()
			if actual != tc.expected {
				t.Fatalf("got: %v; want: %v", actual, tc.expected)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	testCases := []struct {
		name            string
		input, expected Tuple
	}{
		{
			"normalize v(4,0,0)",
			NewVector(4, 0, 0),
			NewVector(1, 0, 0),
		},
		{
			"normalize v(1,2,3)",
			NewVector(1, 2, 3),
			NewVector(0.26726, 0.53452, 0.80178),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Normalize(tc.input)
			if !actual.IsEqual(tc.expected) {
				t.Fatalf("got: %v; want: %v", actual, tc.expected)
			}
		})
	}

	// magnitude of a normalized vector
	actual := Normalize(NewVector(1, 2, 3)).Mag()
	expected := 1.0

	if actual != expected {
		t.Fatalf("got: %v; want: %v", actual, expected)
	}
}

func TestDot(t *testing.T) {
	actual := Dot(NewVector(1, 2, 3), NewVector(2, 3, 4))
	expected := 20.0

	if actual != expected {
		t.Fatalf("got: %v; want: %v", actual, expected)
	}
}

func TestCross(t *testing.T) {
	testCases := []struct {
		name           string
		inputA, inputB Tuple
		expected       Tuple
	}{
		{
			"v(1,2,3) x v(2,3,4)",
			NewVector(1, 2, 3),
			NewVector(2, 3, 4),
			NewVector(-1, 2, -1),
		},
		{
			"v(2,3,4) x v(1,2,3)",
			NewVector(2, 3, 4),
			NewVector(1, 2, 3),
			NewVector(1, -2, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Cross(tc.inputA, tc.inputB)
			if !actual.IsEqual(tc.expected) {
				t.Fatalf("got: %v; want: %v", actual, tc.expected)
			}
		})
	}
}

func TestNewColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	if !equalFloat(c[0], -0.5) {
		t.Errorf("R value is %v. Expected value was -0.5\n", c[0])
	}
	if !equalFloat(c[1], 0.4) {
		t.Errorf("G value is %v. Expected value was 0.4\n", c[1])
	}
	if !equalFloat(c[2], 1.7) {
		t.Errorf("B value is %v. Expected value was 1.7\n", c[2])
	}
}

func TestColorOperations(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1)
	actual, err := Add(c1, c2)
	if err != nil {
		t.Fatalf("error adding colors: %v", err)
	}
	if actual != expected {
		t.Errorf("got: %v; wanted: %v", actual, expected)
	}

	expected = NewColor(0.2, 0.5, 0.5)
	actual, err = Sub(c1, c2)
	for i, val := range actual {
		if !equalFloat(val, expected[i]) {
			t.Errorf("got: %v; wanted: %v", actual, expected)
			break
		}
	}

	color := NewColor(0.2, 0.3, 0.4)
	expected = NewColor(0.4, 0.6, 0.8)
	actual = Scale(color, 2)
	if err != nil {
		t.Fatalf("error multiplying colors: %v", err)
	}
	if actual != expected {
		t.Errorf("got: %v; wanted: %v", actual, expected)
	}
}

func TestHadamardProduct(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)
	actual := HadamardProduct(c1, c2)
	for i, val := range actual {
		if !equalFloat(val, expected[i]) {
			t.Errorf("got: %v; wanted: %v", actual, expected)
			break
		}
	}
}
