package main 

import "testing"

func TestContainer(t *testing.T) {
	c := container{base{42}, "life"}
	got := c.describe()
	againGot := c.define()

	if got != "base with num=42" {
		t.Errorf("Expected base with num=42, got %v", got)
	}

	if againGot != "container with base={42} and str=life" {
		t.Errorf("Expected container with base= num=42 and str=life, got %v", againGot)
	}

}