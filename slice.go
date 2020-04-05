package slice

import (
	"reflect"
	"sort"
)

// SwapFunc represents a function that swaps elements
// at index i and j.
type SwapFunc func(i, j int)

// AutoSwap create SwapFunc from any types of slices.
// It panics when given type is not a slice.
func AutoSwap(slice interface{}) SwapFunc {
	return reflect.Swapper(slice)
}

type stdInterface struct {
	n    int
	swap SwapFunc
	less func(i, j int) bool
}

func (s *stdInterface) Len() int {
	return s.n
}

func (s *stdInterface) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *stdInterface) Swap(i, j int) {
	s.swap(i, j)
}

// Sort sorts elements in slice.
func Sort(n int, swap SwapFunc, less func(i, j int) bool) {
	sort.Sort(&stdInterface{n, swap, less})
}

// MoveLeft moves elements matched function left to the left side of the slice.
// It returns an index of the smallest element not matched left.
// MoveLeft keeps order of the original elements matched left.
func MoveLeft(n int, swap SwapFunc, left func(i int) bool) int {
	var searchFrom int

	for i := 0; i < n; i++ {
		if left(i) {
			if searchFrom <= i {
				searchFrom = i + 1
			}

			swapped := false

			for j := searchFrom; j < n; j++ {
				if !left(j) {
					swap(i, j)
					searchFrom = j + 1
					swapped = true

					break
				}
			}

			if !swapped {
				return i
			}
		}
	}

	return n
}

func Reverse(n int, swap SwapFunc) {
	for i := 0; i < n/2; i++ {
		swap(i, n-1-i)
	}
}
