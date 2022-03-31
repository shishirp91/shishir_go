package main

import (
	"testing"
)

func TestReturnsports(t *testing.T) {
	actualString := Returnsports()
	expectedString := "hockey"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as actual string (%s)", expectedString, actualString)
	}
}
