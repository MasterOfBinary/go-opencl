package opencl

// #cgo CFLAGS: -I${SRCDIR}/../external/include
// #cgo LDFLAGS: ${SRCDIR}/../external/lib/windows/x64/OpenCL.dll
// #include <CL/cl.h>
import "C"
import "unsafe"

type PlatformInfo uint32

const (
	PlatformProfile             PlatformInfo = PlatformInfo(C.CL_PLATFORM_PROFILE)
	PlatformVersion                          = PlatformInfo(C.CL_PLATFORM_VERSION)
	PlatformName                             = PlatformInfo(C.CL_PLATFORM_NAME)
	PlatformVendor                           = PlatformInfo(C.CL_PLATFORM_VENDOR)
	PlatformExtensions                       = PlatformInfo(C.CL_PLATFORM_EXTENSIONS)
	PlatformHostTimerResolution              = PlatformInfo(C.CL_PLATFORM_HOST_TIMER_RESOLUTION)
)

type Platform struct {
	platformID C.cl_platform_id
}

func GetPlatforms() ([]*Platform, error) {
	var n C.cl_uint = C.cl_uint(0)
	errInt := clError(C.clGetPlatformIDs(0, nil, &n))
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	platformIDs := make([]C.cl_platform_id, uint32(n))
	errInt = clError(C.clGetPlatformIDs(n, &platformIDs[0], nil))
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	platforms := make([]*Platform, len(platformIDs))
	for i, platformID := range platformIDs {
		platforms[i] = &Platform{platformID}
	}

	return platforms, nil
}

func (p Platform) GetInfo(name PlatformInfo, output interface{}) error {
	size := uint64(64)
	var sizeRet uint64
	info := make([]byte, size)
	errInt := clError(C.clGetPlatformInfo(p.platformID,
		C.cl_platform_info(name),
		C.size_t(size),
		unsafe.Pointer(&info[0]),
		(*C.size_t)(&sizeRet),
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	// TODO right now just support string pointers
	switch str := output.(type) {
	case *string:
		*str = string(info)
	default:
		return UnexpectedType
	}

	return nil
}
