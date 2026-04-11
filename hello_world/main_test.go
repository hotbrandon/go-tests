package main

import "testing"

func TestHello(t *testing.T) {
	// *testing.T is your hook into the testing framework.
	t.Run("run Hello with Jack in Spanish", func(t *testing.T) {
		expected := "Hola, Jack"
		actual := Hello("Jack", "Spanish")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("run Hello with empty string in Engilsh", func(t *testing.T) {
		expected := "Hello, World"
		actual := Hello("", "English")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("run Hello with empty string in Spanish", func(t *testing.T) {
		expected := "Hola, Mundo"
		actual := Hello("", "Spanish")
		assertCorrectMessage(t, expected, actual)
	})

}

func assertCorrectMessage(t testing.TB, expected, actual string) {
	// testing.TB is an interface that *testing.T and *testing.B both satisfy, so you can use it for both tests and benchmarks.
	t.Helper() // marks this function as a test helper function, so that when it reports an error, the line number will be from the caller of this function, not from inside this function itself.
	if actual != expected {
		t.Errorf("Expected: %q, Got: %q", expected, actual)
	}
}
