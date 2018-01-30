package main

import "testing"

func TestGetArea(t *testing.T) {
	square := square{sideLength: 2}
	triangle := triangle{height: 3, base: 4}

	sa := square.getArea()
	ta := triangle.getArea()

	if sa != 4 {
		t.Errorf("Expected square area to be 4, got %+v", sa)
	}
	if ta != 6 {
		t.Errorf("Expected triangle area to be 3, got %+v", ta)
	}
}
