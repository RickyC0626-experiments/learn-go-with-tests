package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Person", "")
		want := "Hello, Person"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("some dude", spanish)
		want := "Hola, some dude"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("some dude", french)
		want := "Bonjour, some dude"
		assertCorrectMessage(t, got, want)
	})
}
