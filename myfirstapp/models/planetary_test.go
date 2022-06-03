package models_test

import (
	"errors"
	"myfirstapp/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sdhvdsh() error {
	return nil
}

func sdhvdsh2() error {
	return errors.New("sdhjh")
}

func sdhvdsh3() bool {
	return true
}

func sdhvdsh4() bool {
	return false
}

func TestMediaType(t *testing.T) {
	assert.NoError(t, sdhvdsh())
}
func TestMediaType2(t *testing.T) {
	assert.Error(t, sdhvdsh2())
}
func TestMediaType3(t *testing.T) {
	assert.True(t, sdhvdsh3())
}

func TestMediaType4(t *testing.T) {
	assert.False(t, sdhvdsh4())
}

func TestURl(t *testing.T) {
	planetary := models.GetPlanetary()
	assert.Contains(t, planetary.URL, "https://apod.nasa.gov")
}
