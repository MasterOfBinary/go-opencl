package opencl

// #include "opencl.h"
import "C"
import (
	"strings"
	"unsafe"
)

type Program struct {
	program C.cl_program
}

func createProgramWithSource(context Context, programCode string) (Program, error) {
	cs := C.CString(programCode)
	defer C.free(unsafe.Pointer(cs))

	var errInt clError
	program := C.clCreateProgramWithSource(
		context.context,
		1,
		&cs,
		nil,
		(*C.cl_int)(&errInt),
	)
	if errInt != clSuccess {
		return Program{}, clErrorToError(errInt)
	}

	return Program{program}, nil
}

func (p Program) Build(device Device, log *string) error {
	emptyString := C.CString("\x00")
	defer C.free(unsafe.Pointer(emptyString))

	var errInt clError = clError(C.clBuildProgram(
		p.program,
		0,
		nil,
		emptyString,
		nil,
		nil,
	))
	if errInt == clSuccess {
		return nil
	}

	// If there was a log provided, get the compiler log. Otherwise just return the error
	if log == nil {
		return clErrorToError(errInt)
	}

	size := uint64(4096)
	compilerLog := make([]byte, size)
	C.clGetProgramBuildInfo(
		p.program,
		device.deviceID,
		C.CL_PROGRAM_BUILD_LOG,
		C.size_t(size),
		unsafe.Pointer(&compilerLog[0]),
		nil)

	*log = strings.TrimRight(string(compilerLog), "\x00")

	return clErrorToError(errInt)
}

func (p Program) Release() {
	C.clReleaseProgram(p.program)
}

func (p Program) CreateKernel(kernelName string) (Kernel, error) {
	return createKernel(p, kernelName)
}
