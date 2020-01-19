package main

import "testing"

import "fmt"

func TestIteration(t *testing.T) {
	iterate := Iterate("a", 10)
	expected := "aaaaaaaaaa"

	if iterate != expected {
		t.Errorf("got %q but expected %q", iterate, expected)
	}
}

func ExampleIterate() {
	result := Iterate("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 10)
	}
}
