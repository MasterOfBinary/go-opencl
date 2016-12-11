package opencl

// #cgo LDFLAGS: -llibOpenCL
// #include <CL/cl.h>
import "C"

type Platform struct {
}

func (p Platform) GetNumPlatforms() (uint32, error) {
	var n C.cl_uint = C.cl_uint(0)
	errInt := C.clGetPlatformIDs(0, NULL, &n)
	if errInt != C.CL_SUCCESS {
		// TODO handle errors
		panic("Oh no, error")
	}

	return uint32(n)
}
