package opencl

// #cgo !darwin CFLAGS: -I${SRCDIR}/external/include
// #cgo !darwin LDFLAGS: -L${SRCDIR}/external/lib -lOpenCL
// #cgo darwin LDFLAGS: -framework OpenCL
import "C"
