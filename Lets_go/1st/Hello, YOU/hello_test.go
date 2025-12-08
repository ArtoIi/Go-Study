package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Art")
	want := "Hello, Art"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
