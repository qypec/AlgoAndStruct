package heap

import (
	"github.com/stretchr/testify/require"
	"testing"
	)

func TestInit(t *testing.T) {
	var h MinHeap

	h.Init()
	require.Equal(t, h.size, 0)
	require.Equal(t, len(h.arr), 1)
	require.Equal(t, h.arr[0], HeapElement{0, nil})
}
