kernel void kern(global float* out)
{
	size_t i = get_global_id(0);
	out[i] = i;
}