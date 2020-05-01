package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Carlo", "Spanish")
		want := "Hola, Carlo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Renaud", "French")
		want := "Bonjour, Renaud"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in German", func(t *testing.T) {
		got := Hello("Eric", "German")
		want := "Hallo, Eric"

		assertCorrectMessage(t, got, want)
	})
}