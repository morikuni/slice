package slice

import (
	"fmt"
	"testing"

	"github.com/morikuni/slice/internal/assert"
)

func TestMoveLeft(t *testing.T) {
	cases := map[string]struct {
		gen func() ([]int, func(i, j int), func(i int) bool)

		want []int
	}{
		"normal": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

				return is,
					func(i, j int) { is[i], is[j] = is[j], is[i] },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{1, 3, 5, 7, 9},
		},
		"no swap": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}

				return is,
					func(i, j int) { panic("don't swap") },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{1, 3, 5, 7, 9},
		},
		"swap all": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}

				return is,
					func(i, j int) { is[i], is[j] = is[j], is[i] },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{1, 3, 5, 7, 9},
		},
		"size 1 ok": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{1}

				return is,
					func(i, j int) { is[i], is[j] = is[j], is[i] },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{1},
		},
		"size 1 ng": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{2}

				return is,
					func(i, j int) { is[i], is[j] = is[j], is[i] },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{},
		},
		"empty": {
			gen: func() ([]int, func(i, j int), func(i int) bool) {
				is := []int{}

				return is,
					func(i, j int) { is[i], is[j] = is[j], is[i] },
					func(i int) bool { return is[i]%2 == 0 }
			},

			want: []int{},
		},
	}

	for name, tc := range cases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			is, swap, left := tc.gen()
			got := is[:MoveLeft(len(is), swap, left)]

			assert.Equal(t, tc.want, got)
		})
	}
}

func bench(b *testing.B, f func(b *testing.B, slice []int)) {
	size := []int{10, 100, 1000, 10000}

	for _, s := range size {
		s := s
		slice := make([]int, s)

		for i := 0; i < s; i++ {
			slice[i] = i + 1
		}

		b.Run(fmt.Sprint(s), func(b *testing.B) {
			cp := make([]int, s)
			for i := 0; i < b.N; i++ {
				copy(cp, slice)
				f(b, slice)
			}
		})
	}
}

func BenchmarkMoveLeft(b *testing.B) {
	bench(b, func(b *testing.B, slice []int) {
		x := slice[:MoveLeft(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }, func(i int) bool {
			return slice[i]%2 == 1
		})]
		if want, got := len(slice), x[len(x)-1]; want != got {
			b.Errorf("last element mismatch: want=%d got=%d", want, got)
		}
	})
}

func BenchmarkSort(b *testing.B) {
	bench(b, func(b *testing.B, slice []int) {
		Sort(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }, func(i, j int) bool {
			return slice[i] > slice[j]
		})
		if want, got := 1, slice[len(slice)-1]; want != got {
			b.Errorf("last element mismatch: want=%d got=%d", want, got)
		}
	})
}

func BenchmarkAutoSwap(b *testing.B) {
	bench(b, func(b *testing.B, slice []int) {
		Sort(len(slice), AutoSwap(slice), func(i, j int) bool {
			return slice[i] > slice[j]
		})
		if want, got := 1, slice[len(slice)-1]; want != got {
			b.Errorf("last element mismatch: want=%d got=%d", want, got)
		}
	})
}

func TestReverse(t *testing.T) {
	cases := map[string]struct {
		s []int

		want []int
	}{
		"size odd": {
			s:    []int{1, 5, 2, 4, 3},
			want: []int{3, 4, 2, 5, 1},
		},
		"size even": {
			s:    []int{1, 5, 2, 6, 4, 3},
			want: []int{3, 4, 6, 2, 5, 1},
		},
		"size 1": {
			s:    []int{1},
			want: []int{1},
		},
		"size 0": {
			s:    []int{},
			want: []int{},
		},
	}

	for name, tc := range cases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			s := tc.s
			Reverse(len(s), func(i, j int) {
				s[i], s[j] = s[j], s[i]
			})
			assert.Equal(t, tc.want, s)
		})
	}
}
