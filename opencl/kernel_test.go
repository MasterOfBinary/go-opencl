package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateKernel(t *testing.T) {
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

	_, err := program.CreateKernel("main")
	assert.Nil(t, err, "err")
}
