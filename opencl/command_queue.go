package opencl

// #include "opencl.h"
import "C"
import (
	"errors"
	"unsafe"
)

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

func (c CommandQueue) EnqueueNDRangeKernel(kernel *Kernel, workDim uint32, globalWorkSize []uint64) error {
	errInt := clError(C.clEnqueueNDRangeKernel(c.commandQueue,
		kernel.kernel,
		C.cl_uint(workDim),
		nil,
		(*C.size_t)(&globalWorkSize[0]),
		nil, 0, nil, nil))
	return clErrorToError(errInt)
}

func (c CommandQueue) EnqueueReadBuffer(buffer *Buffer, blockingRead bool, dataPtr interface{}) error {
	var br C.cl_bool
	if blockingRead {
		br = C.CL_TRUE
	} else {
		br = C.CL_FALSE
	}

	var ptr unsafe.Pointer
	var dataLen uint64
	switch p := dataPtr.(type) {
	case []float32:
		dataLen = uint64(len(p) * 4)
		ptr = unsafe.Pointer(&p[0])
	default:
		return errors.New("Unexpected type for dataPtr")
	}

	errInt := clError(C.clEnqueueReadBuffer(c.commandQueue,
		buffer.buffer,
		br,
		0,
		C.size_t(dataLen),
		ptr,
		0, nil, nil))
	return clErrorToError(errInt)
}

func (c *CommandQueue) Release() {
	C.clReleaseCommandQueue(c.commandQueue)
}

func (c *CommandQueue) Flush() {
	C.clFlush(c.commandQueue)
}

func (c *CommandQueue) Finish() {
	C.clFinish(c.commandQueue)
}
