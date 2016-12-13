package opencl

// #cgo LDFLAGS: ${SRCDIR}/../external/lib/windows/x64/OpenCL.dll
// #include <CL/cl.h>
import "C"
import (
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
