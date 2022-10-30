package Combination

//组合总数3 https://leetcode.cn/problems/combination-sum-iii/

func combinationSum3(k int, n int) (ans [][]int) {
	path := []int{}
	var pathSum int

	var backTrack func(index int)
	backTrack = func(index int) {
		//剪枝
		if pathSum > n {
			return
		}

		//结束条件
		if len(path) == k {
			if pathSum == n {
				ans = append(ans, append([]int(nil), path...))
			}
			return
		}

		//遍历
		for i := index; i <= 9; i++ {
			path = append(path, i)
			pathSum = pathSum + i
			backTrack(i + 1)
			pathSum = pathSum - i
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return
}
