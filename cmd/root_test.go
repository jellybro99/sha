package cmd

import (
	"os"
	"testing"
)

func TestGetInputsWithArgs(t *testing.T) {
	got, err := getInputs([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0] != "foo" || got[1] != "bar" {
		t.Fatalf("got %v, want [foo bar]", got)
	}
}

func TestGetInputsFromStdin(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = origStdin }()

	go func() {
		w.WriteString("piped message")
		w.Close()
	}()

	got, err := getInputs(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0] != "piped message" {
		t.Fatalf("got %v, want [piped message]", got)
	}
}

func TestGetInputsFromEmptyStdin(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = origStdin }()

	w.Close()

	_, err = getInputs(nil)
	if err == nil {
		t.Fatal("expected error for empty stdin, got nil")
	}
}
