package examples

import "github.com/MasterOfBinary/go-opencl/opencl"

// getFirstDevice returns the first available OpenCL device of type deviceType.
func getFirstDevice(deviceType opencl.DeviceType) opencl.Device {
	platforms, err := opencl.GetPlatforms()
	if err != nil {
		panic(err)
	}

	for _, platform := range platforms {
		var devices []opencl.Device
		devices, err = platform.GetDevices(deviceType)
		if err != nil {
			panic(err)
		}

		for _, device := range devices {
			var available bool
			err = device.GetInfo(opencl.DeviceAvailable, &available)
			if err == nil && available {
				return device
			}
		}
	}

	panic("No device found")
	return opencl.Device{}
}
