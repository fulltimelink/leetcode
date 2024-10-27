package main

func Backward[E any](s []E) func(yield func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

type seq0 func(yield func() bool)

func iter0[Slice ~[]E, E any](s Slice) seq0 {
	return func(yield func() bool) {
		for range s {
			if !yield() {
				return
			}
		}
	}
}
