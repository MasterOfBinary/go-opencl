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

func TestGetDeviceInfo(t *testing.T) {
	p, _ := GetPlatforms()
	d, _ := p[0].GetDevices(DeviceTypeAll)

	var addressBits uint32
	err := d[0].GetInfo(DeviceAddressBits, &addressBits)
	assert.Nil(t, err)
	assert.NotZero(t, addressBits, "device address bits")

	err = d[0].GetInfo(DeviceAddressBits, addressBits)
	assert.NotNil(t, err)

	var deviceAvailable bool
	err = d[0].GetInfo(DeviceAvailable, &deviceAvailable)
	assert.Nil(t, err)

	var bik string
	err = d[0].GetInfo(DeviceBuiltInKernels, &bik)
	assert.Nil(t, err)

	var bik2 []string
	err = d[0].GetInfo(DeviceBuiltInKernels, &bik2)
	assert.Nil(t, err)
}
