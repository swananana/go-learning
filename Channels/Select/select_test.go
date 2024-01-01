// main_test.go

package main

import (
	"testing"
	"time"
)

func TestChannelCommunication(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	select {
	case msg1 := <-c1:
		if msg1 != "one" {
			t.Errorf("Expected 'one', but got %s", msg1)
		}
	case msg2 := <-c2:
		if msg2 != "two" {
			t.Errorf("Expected 'two', but got %s", msg2)
		}
	case <-time.After(3 * time.Second):
		t.Error("Timeout: No message received within 3 seconds")
	}
}
