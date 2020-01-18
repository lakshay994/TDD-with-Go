package main

import "testing"

func TestElaborate(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("returns a string with 'Hello' prepended to the first argument when the second argument is 'english'", func(t *testing.T) {
		got := PrintString("World", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns a string with 'Namaste' prepended to the first argument when the second argument is 'hindi'", func(t *testing.T) {
		got := PrintString("World", "hindi")
		want := "Namaste, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns a string with 'Sat Sri Akal' prepended to the first argument when the second argument is 'punjabi'", func(t *testing.T) {
		got := PrintString("World", "punjabi")
		want := "Sat Sri Akal, World"

		assertCorrectMessage(t, got, want)
	})
}
