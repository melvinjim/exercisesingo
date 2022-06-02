package models_test

import (
	"getangle/models"
	"testing"
)

func TestOperations(t *testing.T) {
	result := models.Operation(12, 7)

	expected := 321.5
	if result != expected {
		t.Errorf("Expected: %f, result: %f", expected, result)
	}
}

func TestOperationsSecond(t *testing.T) {
	result := models.Operation(7, 15)

	expected := 127.5
	if result != expected {
		t.Errorf("Expected: %f, result: %f", expected, result)
	}
}

func TestOperationsThird(t *testing.T) {
	result := models.Operation(2, 39)

	expected := 154.5
	if result != expected {
		t.Errorf("Expected: %f, result: %f", expected, result)
	}
}
