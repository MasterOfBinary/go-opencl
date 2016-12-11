package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-opencl/opencl"
)

func main() {
	n := opencl.GetNumPlatforms()
	fmt.Printf("N: %v\n", n)
}
