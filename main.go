package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-opencl/opencl"
)

const programCode = `
kernel void main(global uchar* in, global uchar* out)
{
	size_t i = get_global_id(0);
	out[i] = in[i] - 3;
}
`

func main() {
	platforms, err := opencl.GetPlatforms()
	if err != nil {
		panic(err)
	}

	var cpuDevice *opencl.Device

	var name string
	for _, platform := range platforms {
		err = platform.GetInfo(opencl.PlatformName, &name)
		if err != nil {
			panic(err)
		}
		var devices []*opencl.Device
		devices, err = platform.GetDevices(opencl.DeviceTypeGPU)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Platform name: %v, number of CPU devices: %v\n", name, len(devices))

		// Use the first device
		if len(devices) > 0 && cpuDevice == nil {
			cpuDevice = devices[0]
		}
	}

	var context *opencl.Context
	context, err = cpuDevice.CreateContext()
	if err != nil {
		panic(err)
	}
	defer context.Release()

	var commandQueue *opencl.CommandQueue
	commandQueue, err = context.CreateCommandQueue()
	if err != nil {
		panic(err)
	}
	defer commandQueue.Release()

	var program *opencl.Program
	program, err = context.CreateProgramWithSource(programCode)
	if err != nil {
		panic(err)
	}
	defer program.Release()

	var log string
	err = program.Build(&log)
	if err != nil {
		fmt.Println(log)
		panic(err)
	}

	_, err = program.CreateKernel("main")
	if err != nil {
		panic(err)
	}
}
