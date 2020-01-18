package main

import "testing"

func TestElaborate(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("returns a string with 'Hello' prepended to the argument string", func(t *testing.T) {
		got := HelloString("Lakshay")
		want := "Hello Lakshay"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns string 'Hello World' if no argument is passed", func(t *testing.T) {
		got := HelloString("")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}
