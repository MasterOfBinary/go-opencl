package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlatforms(t *testing.T) {
	p, err := GetPlatforms()
	assert.Nil(t, err)
	assert.NotEmpty(t, p, "number of platforms")
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

	var ver PlatformMajorMinor
	err = p[0].GetInfo(PlatformVersion, &ver)
	assert.Nil(t, err)

	var exts []string
	err = p[0].GetInfo(PlatformExtensions, &exts)
	assert.Nil(t, err)
}
