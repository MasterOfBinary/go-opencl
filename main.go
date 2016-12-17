package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-opencl/opencl"
)

const (
	dataSize = 128

	programCode = `
kernel void kern(global float* out)
{
	size_t i = get_global_id(0);
	out[i] = i;
}
`
)

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

		var devices []opencl.Device
		devices, err = platform.GetDevices(opencl.DeviceTypeAll)
		if err != nil {
			panic(err)
		}

		version := platform.GetVersion()

		fmt.Printf("Platform name: %v, number of devices: %v, version: %v\n", name, len(devices), version)

		// Use the first device
		if len(devices) > 0 && cpuDevice == nil {
			cpuDevice = &devices[0]
		}
	}

	if cpuDevice == nil {
		panic("No device found")
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

	kernel, err := program.CreateKernel("kern")
	if err != nil {
		panic(err)
	}
	defer kernel.Release()

	buffer, err := context.CreateBuffer([]opencl.MemFlags{opencl.MemWriteOnly}, dataSize*4)
	if err != nil {
		panic(err)
	}
	defer buffer.Release()

	err = kernel.SetArg(0, buffer.Size(), buffer)
	if err != nil {
		panic(err)
	}

	err = commandQueue.EnqueueNDRangeKernel(kernel, 1, []uint64{dataSize})
	if err != nil {
		panic(err)
	}

	commandQueue.Flush()
	commandQueue.Finish()

	data := make([]float32, dataSize)

	err = commandQueue.EnqueueReadBuffer(buffer, true, data)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Output")
	fmt.Println("======")
	for _, item := range data {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
