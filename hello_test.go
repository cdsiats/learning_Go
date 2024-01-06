package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Christian")
		want := "Hello, Christian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when there are no args", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Wrold"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
