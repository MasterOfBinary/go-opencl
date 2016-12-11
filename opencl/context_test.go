package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateContext(t *testing.T) {
	p, _ := GetPlatforms()
	d, _ := p[0].GetDevices(DeviceTypeAll)
	_, err := d[0].CreateContext()
	assert.Nil(t, err)
}
