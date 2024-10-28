package array

import "sort"

func merge(intervals [][]int) [][]int {

	// --  @# 处理极端情况
	if len(intervals) <= 1 {
		return intervals
	}
	// --  @# 排升级序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// --  @# 赋初值
	s := &intervals[0][0]
	e := &intervals[0][1]
	var result [][]int
	for _, v := range intervals {
		// --  @# 记录的e值是否大于当前值的s值为区间是否重合的判断条件
		if *e >= v[0] {
			// --  @# 处理区间完全包含的情况
			if *e <= v[1] {
				e = &v[1]
			}
		} else {
			// --  @# 区间结束，记录当前最大的区间，并为记录值赋初始值
			tv := []int{*s, *e}
			result = append(result, tv)
			s = &v[0]
			e = &v[1]
		}
	}
	// --  @# 追加最后的一个区间
	tv := []int{*s, *e}
	result = append(result, tv)

	return result
}
