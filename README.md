# Go OpenCL
[![Build Status](https://travis-ci.org/MasterOfBinary/go-opencl.svg?branch=master)](https://travis-ci.org/MasterOfBinary/go-opencl)

This is a very simple OpenCL wrapper for Go. To download, use `go get github.com/MasterOfBinary/go-opencl`.

You'll need an OpenCL 2.0 library on all platforms except OS X. Download an SDK and copy its
`libOpenCL.a` file to `opencl/external/lib`. I recommend AMD APP SDK 3.0 or later, which can be downloaded
from [here](http://developer.amd.com/tools-and-sdks/opencl-zone/amd-accelerated-parallel-processing-app-sdk).

To run it, make sure you have an SDK from Intel, NVIDIA, or AMD and a compatible
device. Then run with `go run main.go`:

```
Platform name: NVIDIA CUDA, number of devices: 1, version: 1.2
Platform name: AMD Accelerated Parallel Processing, number of devices: 1, version: 2.0
Platform name: Intel(R) OpenCL, number of devices: 1, version: 1.2
Platform name: Experimental OpenCL 2.1 CPU Only Platform, number of devices: 1, version: 2.1

Output
======
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127
```
