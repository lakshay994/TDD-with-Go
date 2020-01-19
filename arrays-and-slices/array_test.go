package main

import "testing"

import "fmt"

func TestArray(t *testing.T) {

	varArray := []int{2, 4, 6}

	checkValues := func(t *testing.T, sum, expected int, arr []int) {
		t.Helper()
		if sum != expected {
			t.Errorf("got %d but expected %d with %v", sum, expected, arr)
		}
	}

	t.Run("calculates the sum of all numbers in array of variable size", func(t *testing.T) {
		sum := Sum(varArray)
		expected := 12

		checkValues(t, sum, expected, varArray)
	})

}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{2, 4, 6, 8})
	}
}
