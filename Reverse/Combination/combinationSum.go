package Combination

import (
	"fmt"
	"sort"
)

//组合总数 https://leetcode.cn/problems/combination-sum/
//https://blog.csdn.net/S_FMX/article/details/116571815
//没有数量要求，可以无限重复，有总和的限制

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	length := len(candidates)
	var dfs func(target, idx int)

	dfs = func(target, idx int) {
		if idx == length {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)
	return
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates) //sum类型+可重复选择的剪枝要排序

	path := []int{}
	var pathSum int
	var backtrack func(target, idx int)

	backtrack = func(target, idx int) {
		//终止条件
		if pathSum > target {
			return
		}

		//达成条件
		if pathSum == target {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		//遍历
		for i := idx; i < len(candidates); i++ { //剪枝优化，增加判断：pathSum + candidates[i] <= target
			fmt.Println("index=", idx, "i=", i)
			if pathSum+candidates[i] <= target {
				pathSum = pathSum + candidates[i]
				path = append(path, candidates[i])
				backtrack(target, i) //当前数可以重复选则
				pathSum = pathSum - candidates[i]
				path = path[:len(path)-1]
			}
		}

	}
	backtrack(target, 0)
	return
}
