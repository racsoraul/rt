package tuple

import "testing"

func TestIsPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	if !(p.X == 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", p.X)
	}
	if !(p.Y == -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", p.Y)
	}
	if !(p.Z == 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", p.Z)
	}
	if !(p.W == 1.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 1\n", p.W)
	}
	if !p.IsPoint() {
		t.Fatalf("tuple is a vector. It was expected to be a point: %v", p)
	}
}

func TestIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	if !(v.X == 4.3) {
		t.Errorf("tuple 'X' value is %v. Expected value was 4.3\n", v.X)
	}
	if !(v.Y == -4.2) {
		t.Errorf("tuple 'Y' value is %v. Expected value was -4.2\n", v.Y)
	}
	if !(v.Z == 3.1) {
		t.Errorf("tuple 'Z' value is %v. Expected value was 3.1\n", v.Z)
	}
	if !(v.W == 0.0) {
		t.Errorf("tuple 'W' value is %v. Expected value was 0\n", v.W)
	}
	if !v.IsVector() {
		t.Fatalf("tuple is a point. It was expected to be a vector: %v", v)
	}
}