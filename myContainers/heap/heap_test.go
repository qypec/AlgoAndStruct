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


/* accessory type */
type myString struct {
	str string
}

func (s myString) Priority() int {
	return len(s.str)
}
/*               */


func TestInsert(t *testing.T) {
	var h MinHeap

/* Test 01 */
	{
		testsInsert := []int{1, 3, 4, 9, 18, -100}
		resultInsertIndexes := []int{-100, 3, 1, 9, 18, 4}

		h.Init()
		for _, x := range testsInsert {
			h.Insert(x)
		}

		i := 1
		for _, elem := range resultInsertIndexes {
			require.Equal(t, h.arr[i].value.Priority(), elem)
			i++
		}

		h.Reset()
	}

/* Test 02 other types */
	{

		testsInsert := []myString{{"thr"}, {"1"}, {"fivee"},{"tw"}, {"four"}, {""}}
		resultInsert := []HeapElement{
			{1, myString{""}}, {2, myString{"tw"}}, {3, myString{"1"}},
			{4, myString{"thr"}}, {5, myString{"four"}}, {6, myString{"fivee"}},
		}

		h.Init()
		for _, x := range testsInsert {
			h.Insert(x)
		}

		i := 1
		for _, elem := range resultInsert {
			require.Equal(t, h.arr[i].index, elem.index)
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
	h.Insert(1)
	require.Equal(t, h.ExtractMin().Priority(), 1)

/* Test 01 */
	h.Insert(1)
	h.Insert(5)
	require.Equal(t, h.ExtractMin().Priority(), 1)

/* Test 02 */
	require.Equal(t, h.ExtractMin().Priority(), 5)

/* Test 03 */
	h.Insert(0)
	h.Insert(15121)
	require.Equal(t, h.ExtractMin().Priority(), 0)


/* Test 04 */
	require.Equal(t, h.ExtractMin().Priority(), 15121)

/* Test 05 */
	h.Insert(100)
	h.Insert(400)
	h.Insert(10)
	h.Insert(170)
	h.Insert(111410)
	h.Insert(-110)
	h.Insert(17410)
	require.Equal(t, h.ExtractMin().Priority(), -110)
	require.Equal(t, h.ExtractMin().Priority(), 10)
	require.Equal(t, h.ExtractMin().Priority(), 100)
	require.Equal(t, h.ExtractMin().Priority(), 170)
	require.Equal(t, h.ExtractMin().Priority(), 400)
	require.Equal(t, h.ExtractMin().Priority(), 17410)
	require.Equal(t, h.ExtractMin().Priority(), 111410)
}
