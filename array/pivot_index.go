package array

// --  @# 寻找数组的中位数
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	total := 0
	for k, v := range nums {
		if total+total+v == sum {
			return k
		}
		total += v
	}
	return -1
}
