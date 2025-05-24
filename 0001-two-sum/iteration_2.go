package twosum

func TwoSumIteration2(nums []int, target int) []int {
	var length = len(nums)
	for i := range length {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
