package models_test

import (
	"randomcolor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	result := models.RandomInt(0, 15)

	assert.Equal(t, result%2, 0)
	assert.Zero(t, result%2)
}
