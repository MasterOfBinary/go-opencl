package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testBuildHelper(t *testing.T, programCode string, expectValid bool) {
	p, _ := GetPlatforms()
	d, _ := p[0].GetDevices(DeviceTypeAll)
	ctx, _ := d[0].CreateContext()
	_, _ = ctx.CreateCommandQueue()

	program, err := ctx.CreateProgramWithSource(programCode)
	assert.Nil(t, err)

	var log string
	err = program.Build(&log)

	if expectValid {
		assert.Nil(t, err, "err")
	} else {
		assert.NotNil(t, err, "err")
	}
}

func TestBuild(t *testing.T) {
	testBuildHelper(t, `
  kernel void main(global uchar* in, global uchar* out)
  {
    size_t i = get_global_id(0);
    out[i] = in[i] - 3;
  }
  `, true)

	testBuildHelper(t, `
kernel void main(global uchar* in, global uchar* out)
{
  size_t i = get_global_id(0)
  out[i] = in[i] - 3;
}
`, false)
}
