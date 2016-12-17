#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#if __APPLE__
  #include <OpenCL/opencl.h>
#else
  #include <CL/cl.h>
#endif