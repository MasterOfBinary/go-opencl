package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKernel(t *testing.T) {
	var memSize = uint64(128)

	p, _ := GetPlatforms()
	d, _ := p[0].GetDevices(DeviceTypeAll)
	ctx, _ := d[0].CreateContext()
	_, _ = ctx.CreateCommandQueue(d[0])

	programCode := `
  kernel void kern(global uchar* in, global uchar* out)
  {
  	size_t i = get_global_id(0);
  	out[i] = in[i] - 3;
  }
  `

	program, _ := ctx.CreateProgramWithSource(programCode)
	_ = program.Build(d[0], nil)

	kernel, err := program.CreateKernel("kern")
	assert.Nil(t, err, "err")

	buf, _ := ctx.CreateBuffer([]MemFlags{MemWriteOnly}, memSize)

	err = kernel.SetArg(0, buf.Size(), buf)
	assert.Nil(t, err, "err")
}
