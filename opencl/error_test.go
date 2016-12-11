package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClErrorToError(t *testing.T) {
	assert.Nil(t, clErrorToError(clSuccess))
	assert.Equal(t, OutOfHostMemory, clErrorToError(clOutOfHostMemory), "clErrorToError(clOutOfHostMemory)")
}
