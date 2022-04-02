package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Person")

	got := buffer.String()
	want := "Hello, Person"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
