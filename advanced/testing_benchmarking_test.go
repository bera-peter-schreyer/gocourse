package main

import "testing"

func Add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected add(2,3) == 5, got %d", result)
	}
}
