package SumTwoNumber

func twoSum(nums []int, target int) []int {
	results := make([]int, 0)

	for idx1, _ := range nums {
		for idx2 := idx1 + 1; idx2 < len(nums); idx2++ {
			if nums[idx1]+nums[idx2] == target {
				results = append(results, idx1, idx2)
			}
		}
	}

	return results
}
