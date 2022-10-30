package Combination

import "sort"

// unique的 组合总和 https://leetcode.cn/problems/combination-sum-ii/
// 可选集合中有重复元素，但结果集中不能有重复的组合

func combinationSumUnique(candidates []int, target int) (ans [][]int) {
	path := []int{}
	pathSum := 0

	sort.Ints(candidates)
	var backTrack func(index int)
	backTrack = func(index int) {
		//终止条件
		if pathSum == target {
			ans = append(ans, append([]int(nil), path...))
		}

		//遍历
		for i := index; i < len(candidates) && candidates[i]+pathSum < target; i++ {
			//同层使用过的元素跳过
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			//组合path
			path = append(path, candidates[i])
			pathSum = pathSum + candidates[i]
			backTrack(i + 1)
			path = path[:len(path)-1]
			pathSum = pathSum - candidates[i]
		}
	}

	backTrack(0)
	return
}
