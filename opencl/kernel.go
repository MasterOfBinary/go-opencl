package opencl

// #include "opencl.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Kernel struct {
	program *Program
	kernel  C.cl_kernel
}

func createKernel(program *Program, kernelName string) (*Kernel, error) {
	kn := C.CString(kernelName)
	defer C.free(unsafe.Pointer(kn))

	var errInt clError
	kernel := C.clCreateKernel(program.program, kn, (*C.cl_int)(&errInt))
	if errInt != clSuccess {
		fmt.Println("Error code", errInt)
		return nil, clErrorToError(errInt)
	}

	return &Kernel{program, kernel}, nil
}

func (k Kernel) SetArg(argIndex uint32, argSize uint64, argValue interface{}) error {
	var argPtr unsafe.Pointer
	switch argValue.(type) {
	case *Buffer:
		argPtr = unsafe.Pointer(argValue.(*Buffer))
	default:
		return errors.New("Unknown type for argValue")
	}

	errInt := clError(C.clSetKernelArg(
		k.kernel,
		C.cl_uint(argIndex),
		C.size_t(argSize),
		argPtr,
	))
	return clErrorToError(errInt)
}

func (k Kernel) Release() {
	C.clReleaseKernel(k.kernel)
}
