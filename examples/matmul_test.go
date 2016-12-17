package examples

import "github.com/MasterOfBinary/go-opencl/opencl"

func doMatrixMultiplication(device opencl.Device, dataSize int) {
	context, err := device.CreateContext()
	if err != nil {
		panic(err)
	}
	defer context.Release()

	var commandQueue opencl.CommandQueue
	commandQueue, err = context.CreateCommandQueue(device)
	if err != nil {
		panic(err)
	}
	defer commandQueue.Release()

	var program opencl.Program
	program, err = context.CreateProgramWithSourceFile("matmul.cl")
	if err != nil {
		panic(err)
	}
	defer program.Release()

	kernel, err := program.CreateKernel("matmul")
	if err != nil {
		panic(err)
	}
	defer kernel.Release()

	bufferA, err := context.CreateBuffer([]opencl.MemFlags{opencl.MemWriteOnly}, dataSize*4)
	if err != nil {
		panic(err)
	}
	defer bufferA.Release()
}

func ExampleMatrixMultiplication() {
	device := getFirstDevice(opencl.DeviceTypeAll)

	doMatrixMultiplication(device, 1024)
}
