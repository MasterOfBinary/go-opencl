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

}

func ExampleMatrixMultiplication() {
	device := getFirstDevice(opencl.DeviceTypeAll)

	doMatrixMultiplication(device, 1024)
}
