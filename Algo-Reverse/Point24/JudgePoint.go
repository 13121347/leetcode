package Point24

import "math"

//24点游戏  https://leetcode.cn/problems/24-game/
//hard
/**
输入: cards = [4, 1, 8, 7]
输出: true
解释: (8-4) * (7-1) = 24
*/

func judgePoint24(nums []int) bool {
	floatNums := make([]float64, len(nums))
	for i := range floatNums {
		floatNums[i] = float64(nums[i])
	}
	return dfs(floatNums)
}

func dfs(nums []float64) bool {
	//返回结果是不是24
	if len(nums) == 1 {
		return math.Abs(nums[0]-24) < 1e-9 //0.000000001
	}
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//挑选两个数
			n1, n2 := nums[i], nums[j]
			//存放计算后的新数+未选的两个数
			newNums := make([]float64, 0, len(nums))
			//把剩下没有挑的两个数和计算后的数放在这个新数组里
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k])
				}
			}
			//递归计算新的数
			flag = flag || dfs(append(newNums, n1+n2)) //加
			flag = flag || dfs(append(newNums, n1-n2)) //减
			flag = flag || dfs(append(newNums, n2-n1)) //减
			flag = flag || dfs(append(newNums, n1*n2)) //乘
			if n1 != 0 {
				flag = flag || dfs(append(newNums, n2/n1)) //除
			}
			if n2 != 0 {
				flag = flag || dfs(append(newNums, n1/n2)) //除
			}
			if flag {
				return true
			}

		}
	}
	return false
}
