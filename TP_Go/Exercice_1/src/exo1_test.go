package main

import "testing"

func TestGetString(t *testing.T) {
	expected := "This is a string"
	real := GetString()
	if real != expected {
		t.Fatalf("Unvalid string : got %s, expected %s", real, expected)
	}
}