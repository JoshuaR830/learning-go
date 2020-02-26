package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello World"
	message := SayHello()
	if message != expected {
		t.Errorf("Message was incorrect, got: %s, expected: %s.", message, expected)
	}
}