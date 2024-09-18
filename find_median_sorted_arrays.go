package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	l := len(nums3)
	if 0 == l {
		return .0
	}
	if 1 == l {
		return float64(nums3[0])
	}
	sort.Ints(nums3)
	mod := l % 2
	if 1 == mod {
		return float64(nums3[l/2])
	}
	n1 := nums3[(l-1)/2]
	n2 := nums3[l/2]
	return float64(n1+n2) / 2

}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if 0 == l {
		return .0
	}
	if 1 == l {
		nums3 := append(nums1, nums2...)
		return float64(nums3[0])
	}

	mod := l % 2
	pos0, pos1 := l/2, l/2
	if 0 == mod {
		pos0 = (l - 1) / 2
	}
	l1 := len(nums1)
	l2 := len(nums2)
	x, y := 0, 0
	r1, r2 := .0, .0
	for i := 0; i <= pos1; i++ {
		val := .0
		if x == l1 {
			val = float64(nums2[y])
			y++
		} else if y == l2 {
			val = float64(nums1[x])
			x++
		} else if nums1[x] <= nums2[y] {
			val = float64(nums1[x])
			x++
		} else {
			val = float64(nums2[y])
			y++
		}

		if i == pos0 {
			r1 = val
		}
		if i == pos1 {
			r2 = val
		}
	}

	return (r1 + r2) / 2

}
