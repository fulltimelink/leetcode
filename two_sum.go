package main

func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i, v := range nums {
		if index, ok := hash[v]; ok {
			return []int{index, i}
		}
		hash[target-nums[i]] = i
	}
	return []int{-1, -1}
}
