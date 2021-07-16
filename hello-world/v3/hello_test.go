package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Puneet")
	want := "Hello, Puneet"

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
