package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-opencl/opencl"
)

func main() {
	platforms, err := opencl.Get()
	if err != nil {
		panic(err)
	}

	var name string
	for _, platform := range platforms {
		err = platform.GetInfo(clplatform.PlatformName, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Platform name:", name)
	}

}
