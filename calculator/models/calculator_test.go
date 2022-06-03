package models_test

import (
	"calculator/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	n := models.New(29, 37)
	result := n.Multiplication()
	expected := 1073

	assert.Equal(t, result, expected)
}

func TestDivision(t *testing.T) {
	n := models.New(29, 37)
	result := n.Division()
	expected := 0
	if result != expected {
		t.Errorf("Division: Expected: %d, result: %d", expected, result)
	}
}

func TestAddition(t *testing.T) {
	n := models.New(29, 37)
	result := n.Addition()
	expected := 66
	if result != expected {
		t.Errorf("Expected: %d, result: %d", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	n := models.New(29, 37)
	result := n.Subtraction()
	expected := -8
	if result != expected {
		t.Errorf("Expected: %d, result: %d", expected, result)
	}
}
