package main

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "hello Chris"

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
