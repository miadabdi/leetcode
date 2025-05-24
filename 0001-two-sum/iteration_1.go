package twosum

func TwoSumIteration1(nums []int, target int) []int {
	for i, ei := range nums {
		for j, ej := range nums {
			if i == j {
				continue
			}

			if ei+ej == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
