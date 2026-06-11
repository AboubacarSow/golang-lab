package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	input := "Hello world Golang community"

	result := cleanInput(input)
	expected := []string{"hello", "world", "golang", "community"}

	if len(result) != len(expected) {
		t.Errorf("Expected: %d but Actual: %d", len(expected), len(result))
	}
	for i, word := range expected {
		if result[i] != word {
			t.Errorf("Expected: %s but Actual:%s", word, result[i])
		}
	}

}
