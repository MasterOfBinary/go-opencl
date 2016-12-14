package opencl

// #cgo CFLAGS: -I${SRCDIR}/../external/include
// #cgo LDFLAGS: -lOpenCL
// #include <CL/cl.h>
import "C"
import (
	"errors"
	"fmt"
)

type clError int32

const (
	clSuccess             clError = clError(C.CL_SUCCESS)
	clDeviceNotFound              = clError(C.CL_DEVICE_NOT_FOUND)
	clBuildProgramFailure         = clError(C.CL_BUILD_PROGRAM_FAILURE)
	clOutOfHostMemory             = clError(C.CL_OUT_OF_HOST_MEMORY)
)

var (
	DeviceNotFound      = errors.New("Device not found")
	BuildProgramFailure = errors.New("Build program failure")
	OutOfHostMemory     = errors.New("Out of host memory")

	UnexpectedType      = errors.New("Unexpected type")
	ErrorParsingVersion = errors.New("Error parsing OpenCL version")
	UnknownError        = errors.New("Unknown error")
)

var (
	errorMap = map[clError]error{
		clSuccess:             nil, // Probably never used
		clDeviceNotFound:      DeviceNotFound,
		clBuildProgramFailure: BuildProgramFailure,
		clOutOfHostMemory:     OutOfHostMemory,
	}
)

func clErrorToError(clerr clError) error {
	err, ok := errorMap[clerr]
	if ok {
		return err
	}
	fmt.Println("Error", clerr)
	return UnknownError
}
