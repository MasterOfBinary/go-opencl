package opencl

// #cgo LDFLAGS: -lOpenCL
// #include <CL/cl.h>
import "C"

type Context struct {
	context C.cl_context
	device  *Device
}

func createContext(device *Device) (*Context, error) {
	// TODO add more functionality. Super simple context creation right now
	var errInt clError
	ctx := C.clCreateContext(
		nil,
		1,
		(*C.cl_device_id)(&device.deviceID),
		nil,
		nil,
		(*C.cl_int)(&errInt),
	)
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	return &Context{ctx, device}, nil
}

func (c *Context) CreateCommandQueue() (*CommandQueue, error) {
	return createCommandQueue(c)
}

func (c *Context) CreateProgramWithSource(programCode string) (*Program, error) {
	return createProgramWithSource(c, programCode)
}

func (c *Context) CreateBuffer(memFlags []MemFlags, size uint64) (*Buffer, error) {
	return createBuffer(c, memFlags, size)
}

func (c *Context) Release() {
	C.clReleaseContext(c.context)
}
