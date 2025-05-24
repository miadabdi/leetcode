package twosum

// Find two numbers in a slice to sum up to target
func TwoSumIteration3(nums []int, target int) []int {
	var existance = map[int]int{}

	for i, el := range nums {
		var complement = target - el
		if val, ok := existance[complement]; ok {
			if val != i {
				return []int{val, i}
			}
		}
		existance[el] = i
	}

	return []int{-1, -1}
}
