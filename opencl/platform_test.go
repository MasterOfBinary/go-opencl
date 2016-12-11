package opencl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlatforms(t *testing.T) {
	p, err := GetPlatforms()
	assert.Nil(t, err)
	assert.NotEmpty(t, p, "number of platforms")
	fmt.Printf("Number: %v\n", len(p))
}

func TestGetInfo(t *testing.T) {
	p, err := GetPlatforms()
	assert.Nil(t, err)
	assert.NotEmpty(t, p, "number of platforms")

	var name string
	err = p[0].GetInfo(PlatformName, &name)
	assert.Nil(t, err)
	assert.NotEmpty(t, name, "platform name")

	err = p[0].GetInfo(PlatformName, name)
	assert.NotNil(t, err)
}
