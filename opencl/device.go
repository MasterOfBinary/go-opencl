package opencl

// #cgo LDFLAGS: -lOpenCL
// #include <CL/cl.h>
import "C"
import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// DeviceType is a type of OpenCL device.
type DeviceType uint32

// DeviceType constants.
const (
	DeviceTypeDefault     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DeviceTypeCPU                    = C.CL_DEVICE_TYPE_CPU
	DeviceTypeGPU                    = C.CL_DEVICE_TYPE_GPU
	DeviceTypeAccelerator            = C.CL_DEVICE_TYPE_ACCELERATOR
	DeviceTypeCustom                 = C.CL_DEVICE_TYPE_CUSTOM
	DeviceTypeAll                    = C.CL_DEVICE_TYPE_ALL
)

// DeviceInfo is a type of info that can be retrieved by Device.GetInfo.
type DeviceInfo uint32

// DeviceInfo constants.
const (
	DeviceAddressBits       DeviceInfo = DeviceInfo(C.CL_DEVICE_ADDRESS_BITS)
	DeviceAvailable                    = DeviceInfo(C.CL_DEVICE_AVAILABLE)
	DeviceBuiltInKernels               = DeviceInfo(C.CL_DEVICE_BUILT_IN_KERNELS)
	DeviceCompilerAvailable            = DeviceInfo(C.CL_DEVICE_COMPILER_AVAILABLE)
)

var (
	deviceInfoTypes = map[DeviceInfo][]interface{}{
		DeviceAddressBits:       {uint32(0)},
		DeviceAvailable:         {false},
		DeviceBuiltInKernels:    {"", []string{}},
		DeviceCompilerAvailable: {false},
	}
)

// Device is a structure for an OpenCL device.
type Device struct {
	deviceID C.cl_device_id
}

// getDevices returns a slice of devices of type deviceType for platform.
func getDevices(platform Platform, deviceType DeviceType) ([]Device, error) {
	var deviceCount C.cl_uint = C.cl_uint(0)
	errInt := clError(C.clGetDeviceIDs(platform.platformID, C.cl_device_type(deviceType), 0, nil, &deviceCount))
	if errInt != clSuccess {
		if errInt == clDeviceNotFound {
			return []Device{}, nil
		}
		return nil, clErrorToError(errInt)
	}

	deviceIDs := make([]C.cl_device_id, uint32(deviceCount))
	errInt = clError(C.clGetDeviceIDs(
		platform.platformID,
		C.cl_device_type(deviceType),
		deviceCount,
		&deviceIDs[0],
		nil,
	))
	if errInt != clSuccess {
		if errInt == clDeviceNotFound {
			return []Device{}, nil
		}
		return nil, clErrorToError(errInt)
	}

	devices := make([]Device, len(deviceIDs))
	for i, deviceID := range deviceIDs {
		devices[i] = Device{deviceID}
	}

	return devices, nil
}

// CreateContext creates and returns an OpenCL context for a device. After use the
// context must be released.
func (d Device) CreateContext() (*Context, error) {
	return createContext(d)
}

// GetInfo retrieves the information specified by name and stores it in output.
// The output must correspond to the return type for that type of info:
//
// DeviceAddressBits *bool
// DeviceAvailable *bool
// DeviceBuiltInKernels *string or *[]string
// DeviceCompilerAvailable *bool
//
// Note that if DeviceBuiltInKernels is retrieved with output being a *string,
// the extensions will be a semicolon-separated list as specified by the OpenCL
// reference for clGetDeviceInfo.
func (d Device) GetInfo(name DeviceInfo, output interface{}) error {
	var size uint64
	errInt := clError(C.clGetDeviceInfo(
		d.deviceID,
		C.cl_device_info(name),
		0,
		nil,
		(*C.size_t)(&size),
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	info := make([]byte, size)
	errInt = clError(C.clGetPlatformInfo(
		d.deviceID,
		C.cl_device_info(name),
		C.size_t(size),
		unsafe.Pointer(&info[0]),
		nil,
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	expectedTypes, ok := deviceInfoTypes[name]
	if !ok {
		return UnexpectedType
	}

	// Output interface needs to be a pointer to the expected type
	v := reflect.ValueOf(output)
	if v.Kind() != reflect.Ptr {
		return UnexpectedType
	}

	elem := v.Elem()

	validType := false
	for _, curType := range expectedTypes {
		if elem.Type() == reflect.TypeOf(curType) {
			validType = true
			break
		}
	}

	if !validType {
		return UnexpectedType
	}

	switch t := output.(type) {
	case *string:
		outputString := zeroTerminatedByteSliceToString(info)
		*t = outputString
	case *uint32:
		*t = *(*uint32)(unsafe.Pointer(&info))
	case *bool:
		u := *(*C.cl_bool)(unsafe.Pointer(&info))
		*t = u == C.CL_TRUE
	case *[]string:
		if name == DeviceBuiltInKernels {
			outputString := zeroTerminatedByteSliceToString(info)
			elems := strings.Split(outputString, ";")
			*t = elems
		}
	}

	return nil
}
