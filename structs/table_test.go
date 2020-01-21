package main

import "testing"

func TestWithTables(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 200.00},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		want := test.want

		if got != want {
			t.Errorf("got %0.2f but got %0.2f", got, want)
		}
	}
}
