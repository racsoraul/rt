package tuple

import "testing"

func TestIsPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	if !equalFloat(p.X, 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", p.X)
	}
	if !equalFloat(p.Y, -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", p.Y)
	}
	if !equalFloat(p.Z, 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", p.Z)
	}
	if !equalFloat(p.W, 1.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 1\n", p.W)
	}
	if !p.IsPoint() {
		t.Fatalf("tuple is a vector. It was expected to be a point: %v", p)
	}
}

func TestIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	if !equalFloat(v.X, 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", v.X)
	}
	if !equalFloat(v.Y, -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", v.Y)
	}
	if !equalFloat(v.Z, 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", v.Z)
	}
	if !equalFloat(v.W, 0.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 0\n", v.W)
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
