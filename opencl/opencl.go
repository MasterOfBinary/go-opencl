// +build !darwin

package opencl

// #cgo CFLAGS: -I${SRCDIR}/external/include
// #cgo LDFLAGS: -L${SRCDIR}/external/lib -lOpenCL
// #include <CL/cl.h>
import "C"
