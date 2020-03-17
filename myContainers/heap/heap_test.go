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
	require.Equal(t, h.arr[0], HeapElement{0, 0,nil})
}

func TestInsert(t *testing.T) {
	var h MinHeap

/* Test 01 */
	{
		testsInsert := []int{1, 3, 4, 9, 18, -100}
		resultInsertIndexes := []int{-100, 3, 1, 9, 18, 4}

		h.Init()
		for _, x := range testsInsert {
			h.Insert(x, x)
		}

		i := 1
		for _, elem := range resultInsertIndexes {
			require.Equal(t, h.arr[i].priority, elem)
			i++
		}

		h.Reset()
	}

/* Test 02 other types */
	{

		testsInsert := []string{"thr", "1", "fivee", "tw", "four", ""}
		resultInsert := []HeapElement{
			{1, 0, ""}, {2, 2, "tw"}, {3, 1, "1"},
			{4, 3, "thr"}, {5, 4, "four"}, {6, 5, "fivee"},
		}

		h.Init()
		for _, x := range testsInsert {
			h.Insert(len(x), x)
		}

		i := 1
		for _, elem := range resultInsert {
			require.Equal(t, h.arr[i].index, elem.index)
			require.Equal(t, h.arr[i].priority, elem.priority)
			require.Equal(t, h.arr[i].value, elem.value)
			i++
		}

		h.Reset()
	}
}

func TestExtractMin(t *testing.T) {
	var h MinHeap

	h.Init()

/* Test 00 basic */
	h.Insert(1, 1)
	require.Equal(t, h.ExtractMin().(int), 1)

/* Test 01 */
	h.Insert(1, 1)
	h.Insert(5, 5)
	require.Equal(t, h.ExtractMin().(int), 1)

/* Test 02 */
	require.Equal(t, h.ExtractMin().(int), 5)

/* Test 03 */
	h.Insert(0, 0)
	h.Insert(15121, 15121)
	require.Equal(t, h.ExtractMin().(int), 0)


/* Test 04 */
	require.Equal(t, h.ExtractMin().(int), 15121)

/* Test 05 */
	h.Insert(100, 100)
	h.Insert(400, 400)
	h.Insert(10, 10)
	h.Insert(170, 170)
	h.Insert(111410, 111410)
	h.Insert(-110, -110)
	h.Insert(17410, 17410)
	require.Equal(t, h.ExtractMin().(int), -110)
	require.Equal(t, h.ExtractMin().(int), 10)
	require.Equal(t, h.ExtractMin().(int), 100)
	require.Equal(t, h.ExtractMin().(int), 170)
	require.Equal(t, h.ExtractMin().(int), 400)
	require.Equal(t, h.ExtractMin().(int), 17410)
	require.Equal(t, h.ExtractMin().(int), 111410)
}
