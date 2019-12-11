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

// Remove moves elements matched function remove to the tail of slice.
// It returns an index of the head element not matched function remove.
// Remove keeps order of the original elements.
func Remove(n int, swap SwapFunc, remove func(i int) bool) int {
	var searchFrom int
	for i := 0; i < n; i++ {
		if remove(i) {
			if searchFrom <= i {
				searchFrom = i + 1
			}
			swapped := false
			for j := searchFrom; j < n; j++ {
				if !remove(j) {
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
