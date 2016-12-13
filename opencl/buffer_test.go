package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBuffer(t *testing.T) {
	memSize := uint64(128)

	p, _ := GetPlatforms()
	d, _ := p[0].GetDevices(DeviceTypeAll)
	ctx, _ := d[0].CreateContext()
	_, err := ctx.CreateBuffer([]MemFlags{MemWriteOnly}, memSize)
	assert.Nil(t, err, "err")
}
