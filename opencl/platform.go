package opencl

// #cgo LDFLAGS: -lOpenCL
// #include <CL/cl.h>
import "C"
import (
	"errors"
	"strconv"
	"strings"
	"unsafe"
)

// PlatformInfo is a type of info that can be retrieved by Platform.GetInfo.
type PlatformInfo uint32

// PlatformInfo constants.
const (
	PlatformProfile    PlatformInfo = PlatformInfo(C.CL_PLATFORM_PROFILE)
	PlatformVersion                 = PlatformInfo(C.CL_PLATFORM_VERSION)
	PlatformName                    = PlatformInfo(C.CL_PLATFORM_NAME)
	PlatformVendor                  = PlatformInfo(C.CL_PLATFORM_VENDOR)
	PlatformExtensions              = PlatformInfo(C.CL_PLATFORM_EXTENSIONS)
)

// PlatformMajorMinor contains the major and minor versions of an OpenCL
// platform.
type PlatformMajorMinor struct {
	Major uint8
	Minor uint8
}

// Platform is a structure for an OpenCL platform.
type Platform struct {
	platformID C.cl_platform_id
	version    *PlatformMajorMinor
}

// GetPlatforms returns a slice containing all platforms available.
func GetPlatforms() ([]*Platform, error) {
	var platformCount C.cl_uint = C.cl_uint(0)
	errInt := clError(C.clGetPlatformIDs(0, nil, &platformCount))
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	platformIDs := make([]C.cl_platform_id, uint32(platformCount))
	errInt = clError(C.clGetPlatformIDs(platformCount, &platformIDs[0], nil))
	if errInt != clSuccess {
		return nil, clErrorToError(errInt)
	}

	platforms := make([]*Platform, len(platformIDs))
	for i, platformID := range platformIDs {
		platforms[i] = &Platform{
			platformID: platformID,
		}

		ver, err := platforms[i].GetVersion()
		if err == nil {
			platforms[i].version = ver
		}
	}

	return platforms, nil
}

// GetInfo retrieves the information specified by name and stores it in output.
// The output must correspond to the return type for that type of info:
//
// PlatformProfile *string
// PlatformVersion *string or *PlatformMajorMinor
// PlatformName *string
// PlatformVendor *string
// PlatformExtensions *[]string or *string
// PlatformICDSuffixKHR *string
//
// Note that if PlatformExtensions is retrieved with output being a *string,
// the extensions will be a space-separated list as specified by the OpenCL
// reference for clGetPlatformInfo.
func (p Platform) GetInfo(name PlatformInfo, output interface{}) error {
	var size uint64
	errInt := clError(C.clGetPlatformInfo(
		p.platformID,
		C.cl_platform_info(name),
		0,
		nil,
		(*C.size_t)(&size),
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	info := make([]byte, size)
	errInt = clError(C.clGetPlatformInfo(
		p.platformID,
		C.cl_platform_info(name),
		C.size_t(size),
		unsafe.Pointer(&info[0]),
		nil,
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	outputString := zeroTerminatedByteSliceToString(info)

	switch t := output.(type) {
	case *string:
		*t = outputString
	case *PlatformMajorMinor:
		if name != PlatformVersion {
			return UnexpectedType
		}

		ver, errVer := parseVersion(outputString)
		if errVer != nil {
			return errors.New("Unable to parse platform info to *PlatformMajorMinor, err: " + errVer.Error())
		}

		*t = *ver

	case *[]string:
		if name != PlatformExtensions {
			return UnexpectedType
		}

		elems := strings.Split(outputString, " ")
		*t = elems

	default:
		return UnexpectedType
	}

	return nil
}

// GetDevices returns a slice of devices of type deviceType for a Platform. If there are
// no such devices it returns an empty slice.
func (p Platform) GetDevices(deviceType DeviceType) ([]Device, error) {
	return getDevices(p, deviceType)
}

// GetVersion returns the platform OpenCL version.
func (p *Platform) GetVersion() (*PlatformMajorMinor, error) {
	var ver PlatformMajorMinor
	err := p.GetInfo(PlatformVersion, &ver)
	if err != nil {
		return nil, err
	}

	p.version = &ver

	return &ver, nil
}

// parseVersion is a helper function to parse an OpenCL version. The version format
// is given by the specification to be:
//
// OpenCL<space><major_version.minor_version><space><platform-specific information>
//
// The only part that concerns us here is the major/minor version combination.
func parseVersion(ver string) (*PlatformMajorMinor, error) {
	elems := strings.SplitN(ver, " ", 3)
	if len(elems) < 3 || elems[0] != "OpenCL" {
		return nil, ErrorParsingVersion
	}

	verElems := strings.Split(elems[1], ".")
	if len(verElems) != 2 {
		return nil, ErrorParsingVersion
	}

	maj, errMaj := strconv.ParseUint(verElems[0], 10, 64)
	if errMaj != nil {
		return nil, ErrorParsingVersion
	}
	min, errMin := strconv.ParseUint(verElems[1], 10, 64)
	if errMin != nil {
		return nil, ErrorParsingVersion
	}

	return &PlatformMajorMinor{
		Major: uint8(maj),
		Minor: uint8(min),
	}, nil
}
