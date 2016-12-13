package opencl

// TODO this is a cheat because my OpenCL DLL is too old to support
// clCreateCommandQueueWithProperties but in the headers I'm using it's deprecated,
// causing a Go error

// #cgo LDFLAGS: -lOpenCL
// #pragma GCC diagnostic ignored "-Wdeprecated-declarations"
// #include <CL/cl.h>
import "C"

type CommandQueue struct {
	commandQueue C.cl_command_queue
	context      *Context
}

func createCommandQueue(context *Context) (*CommandQueue, error) {
	var errInt clError
	queue := C.clCreateCommandQueue(
		context.context,
		context.device.deviceID,
		0,
		(*C.cl_int)(&errInt),
	)
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	return &CommandQueue{queue, context}, nil
}

func (c *CommandQueue) Release() {
	C.clReleaseCommandQueue(c.commandQueue)
}
