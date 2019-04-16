package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Boris")
	want := "Hello, Boris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}