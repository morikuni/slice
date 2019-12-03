package slice

import (
	"reflect"
	"sort"
)

type Swap func(i, j int)

func AutoSwap(slice interface{}) Swap {
	return reflect.Swapper(slice)
}

type stdInterface struct {
	n    int
	swap func(i, j int)
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

func Sort(n int, swap Swap, less func(i, j int) bool) {
	sort.Sort(&stdInterface{n, swap, less})
}

func Remove(n int, swap Swap, remove func(i int) bool) int {
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
