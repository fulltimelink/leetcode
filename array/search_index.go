package array

// --  @# 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// --  @# 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int {
	l := len(nums)
	if 0 == l || nums[0] >= target {
		return 0
	}
	left := 0
	right := l - 1
	for left < right {
		mid := left + (right-left)/2 + (right-left)%2
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	index := left
	if index > right {
		index = right
	}
	if nums[index] < target {
		return index + 1
	}
	return index

}
