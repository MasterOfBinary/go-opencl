package opencl

import "strings"

// zeroTerminatedByteSliceToString converts input to a string and trims the
// bytes after the terminating zero.
func zeroTerminatedByteSliceToString(input []byte) string {
	return strings.TrimRight(string(input), "\x00")
}
