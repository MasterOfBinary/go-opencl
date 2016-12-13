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
	_, _ = ctx.CreateCommandQueue()

	programCode := `
  kernel void main(global uchar* in, global uchar* out)
  {
  	size_t i = get_global_id(0);
  	out[i] = in[i] - 3;
  }
  `

	program, _ := ctx.CreateProgramWithSource(programCode)
	_ = program.Build(nil)

	_, _ = program.CreateKernel("main")

	_, err := ctx.CreateBuffer([]MemFlags{MemWriteOnly}, memSize)
	assert.Nil(t, err, "err")
}
