package models_test

import (
	"calculatorofconsol/models"
	"testing"
)

func TestAdd(t *testing.T) {
	result := models.Add()
	expected := 0
	if result != expected {
		t.Errorf(" Expected: %d, result: %d", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := models.Subtraction()
	expected := 0
	if result != expected {
		t.Errorf(" Expected: %d, result: %d", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	result := models.Multiplication()
	expected := 0
	if result != expected {
		t.Errorf(" Expected: %d, result: %d", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result := models.Division(10, 2)
	expected := 5
	if result != expected {
		t.Errorf(" Expected: %d, result: %d", expected, result)
	}
}
