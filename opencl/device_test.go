package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevices(t *testing.T) {
	p, _ := GetPlatforms()
	d, err := p[0].GetDevices(DeviceTypeAll)
	assert.Nil(t, err)
	assert.NotEmpty(t, d, "number of devices")
}

/*func TestGetInfo(t *testing.T) {
	p, err := GetPlatforms()
	assert.Nil(t, err)
	assert.NotEmpty(t, p, "number of platforms")

	var name string
	err = p[0].GetInfo(PlatformName, &name)
	assert.Nil(t, err)
	assert.NotEmpty(t, name, "platform name")

	err = p[0].GetInfo(PlatformName, name)
	assert.NotNil(t, err)
}*/
