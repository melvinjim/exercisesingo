package models_test

import (
	"github.com/stretchr/testify/assert"
	"randomcolor/models"
	"testing"
)

func TestNumber(t *testing.T) {
	result := models.RandomInt(0, 15)

	assert.Equal(t, result%2, 0)
}
