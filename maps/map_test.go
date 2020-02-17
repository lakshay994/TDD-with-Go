package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	t.Run("search dictionary", func(t *testing.T) {
		got, _ := dictionary.Search(dictionary, "test")
		want := "just a test"

		assertStrings(t, got, want)
	})

	t.Run("word does not exist", func(t *testing.T) {
		_, err := dictionary.Search(dictionary, "random")
		want := ErrorResponse

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, received, expected string) {
	t.Helper()

	if received != expected {
		t.Errorf("got %q but wanted %q", received, expected)
	}
}
