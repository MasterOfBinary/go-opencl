/*
 * matmul.cl
 *
 * Perform matrix multiplication.
 */

__kernel void matmul(
          __global float* outC,
          __global float* inA,
          __global float* inB,
          int widthA, int widthB)
{
   int tx = get_global_id(0);
   int ty = get_global_id(1);

   float value = 0;
   for (int i = 0; i < widthA; i++)
   {
      float curA = inA[ty * widthA + i];
      float curB = inB[i * widthB + tx];
      value += curA * curB;
   }

   outC[ty * widthA + tx] = value;
}