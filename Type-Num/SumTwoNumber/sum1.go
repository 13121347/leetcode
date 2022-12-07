package SumTwoNumber

// 暴力解法，时间复杂度O(N2)
func TwoSum(nums []int, target int) []int {
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

// Hash解法，时间复杂度O（N）
func TwoSumHash(nums []int, target int) []int {
	results := make([]int, 0)
	numsMapping := make(map[int]int)

	for idx, value := range nums {
		numsMapping[value] = idx
		targetIdx, ok := numsMapping[target-value]
		//最后的结果需要去重，这里golang不提供类似java的hashmap方法,但是题干上说只有一组结果，所以这里在ok的时候就可以直接返回了
		if ok {
			results = append(results, idx, targetIdx)
			return results
		}
	}
	return nil
}
