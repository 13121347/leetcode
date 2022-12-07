package Algo_Dynamic

/**
* 同样是机器人走网格，这次需要选一个经过路径上数字和最小，我想的是把所有路径都找出来，然后把值相加，找到和最小的
* 但是这个动态规划给的解法是，到dp[i][j] =  min(dp[i-1][j]，dp[i][j-1]) + arr[i][j]
* 问题是，如果按照动态规划来解，下面的路径找到的就不是最小的了，路径会是 1-> 1-> 1000-> 1-> 1; 但是最小路径明明是1-> 3-> 1-> 1-> 1
* [1,3,1],
* [1,10001,1],
* [1000,1,1]
* 不对！ 这里一定会找到最小的，因为每个格子都是用前两个格子的min来算的！ 这里已经做了一次抉择
 */

func lazyRobot(c, r int) int {
	arr := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	if c <= 0 || r <= 0 {
		return 0
	}
	dp := [100][100]int{}
	for i := 0; i < c; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < r; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < c; i++ {
		for j := 1; j < r; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + arr[i][j]
		}
	}
	return dp[c-1][r-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
