package main

import "testing"

func TestStructs(t *testing.T) {

	validateRectanglePerimeter := func(t *testing.T, got, want float64, rectangle Rectangle) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f but expected %.2f for rectangle dimensions %v", got, want, rectangle)
		}
	}

	validateArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f but expected %.2f", got, want)
		}
	}

	t.Run("perimeter function for Rectangle struct and check it against expected o/p", func(t *testing.T) {
		rectangle := Rectangle{length: 20.0, breadth: 10.0}
		perimeter := rectangle.Perimeter()
		expected := 60.00

		validateRectanglePerimeter(t, perimeter, expected, rectangle)
	})

	t.Run("area function for Rectangle struct and check it against expected o/p", func(t *testing.T) {
		rectangle := Rectangle{length: 20.0, breadth: 10.0}
		expected := 200.00

		validateArea(t, rectangle, expected)
	})

	t.Run("area function for Circle struct and check it against expected o/p", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		expected := 314.1592653589793

		validateArea(t, circle, expected)
	})

}
