package opencl

// #cgo CFLAGS: -I${SRCDIR}/../external/include
// #cgo LDFLAGS: ${SRCDIR}/../external/lib/windows/x64/OpenCL.dll
// #include <CL/cl.h>
import "C"

type DeviceType uint32

const (
	DeviceTypeDefault     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DeviceTypeCPU                    = C.CL_DEVICE_TYPE_CPU
	DeviceTypeGPU                    = C.CL_DEVICE_TYPE_GPU
	DeviceTypeAccelerator            = C.CL_DEVICE_TYPE_ACCELERATOR
	DeviceTypeCustom                 = C.CL_DEVICE_TYPE_CUSTOM
	DeviceTypeAll                    = C.CL_DEVICE_TYPE_ALL
)

type Device struct {
	deviceID C.cl_device_id
}

func getDevices(platformID C.cl_platform_id, deviceType DeviceType) ([]*Device, error) {
	var deviceCount C.cl_uint = C.cl_uint(0)
	errInt := clError(C.clGetDeviceIDs(platformID, C.cl_device_type(deviceType), 0, nil, &deviceCount))
	if errInt != clSuccess {
		if errInt == clDeviceNotFound {
			return []*Device{}, nil
		}
		return nil, clErrorToError(errInt)
	}

	deviceIDs := make([]C.cl_device_id, uint32(deviceCount))
	errInt = clError(C.clGetDeviceIDs(platformID, C.cl_device_type(deviceType), deviceCount, &deviceIDs[0], nil))
	if errInt != clSuccess {
		if errInt == clDeviceNotFound {
			return []*Device{}, nil
		}
		return nil, clErrorToError(errInt)
	}

	devices := make([]*Device, len(deviceIDs))
	for i, deviceID := range deviceIDs {
		devices[i] = &Device{deviceID}
	}

	return devices, nil
}

func (d *Device) CreateContext() (*Context, error) {
	return createContext(d)
}
