package leetcode

// Problem 001 - Two Sum , Solution #1
func TwoSumSolution1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Problem 001 - Two Sum , Solution #2
func TwoSumSolution2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
