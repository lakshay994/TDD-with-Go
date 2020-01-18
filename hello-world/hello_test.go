package main

import "testing"

func TestHello(t *testing.T) {
	got := HelloString("World")
	want := "Hello World"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
