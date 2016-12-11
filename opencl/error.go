package opencl

// #cgo CFLAGS: -I${SRCDIR}/../external/include
// #cgo LDFLAGS: ${SRCDIR}/../external/lib/windows/x64/OpenCL.dll
// #include <CL/cl.h>
import "C"
import "errors"

type clError int32

const (
	clSuccess         clError = clError(C.CL_SUCCESS)
	clOutOfHostMemory         = clError(C.CL_OUT_OF_HOST_MEMORY)
)

var (
	OutOfHostMemory = errors.New("Out of host memory")

	UnexpectedType = errors.New("Unexpected type")
	UnknownError   = errors.New("Unknown error")
)

var (
	errorMap = map[clError]error{
		clSuccess:         nil, // Probably never used
		clOutOfHostMemory: OutOfHostMemory,
	}
)

func clErrorToError(clerr clError) error {
	err, ok := errorMap[clerr]
	if ok {
		return err
	}
	return UnknownError
}
