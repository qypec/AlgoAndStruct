package binary_search

func ascendingMoves(a []int, m int, toFind int, fix, l, r *int) {
	if a[m] >= toFind {
		if a[m] == toFind {
			*fix = m
		}
		*r = m - 1
	} else {
		*l = m + 1
	}
}

func descendingMoves(a []int, m int, toFind int, fix, l, r *int) {
	if a[m] > toFind {
		*l = m + 1
	} else {
		if a[m] == toFind {
			*fix = m
		}
		*r = m - 1
	}
}

/* finds the first element that is equal than the given (toFind) */
/* array a must be sorted */
/* Complexity: log(n) */
func BinarySearch(a []int, toFind int, l, r int) int {
	moves := ascendingMoves
	if a[0] > a[len(a)-1] {
		moves = descendingMoves
	}

	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		moves(a, m, toFind, &fix, &l, &r)
	}
	return fix
}
