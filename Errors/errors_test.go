package main

import "testing"

func TestF1(t *testing.T) {
	_, got := f1(42)

	if got == nil {
		t.Error("Expected an error")
	}

	if got.Error() != "can't work with 42" {
		t.Error("Expected can't work with 42")
	}
}