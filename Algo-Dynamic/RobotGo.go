package Algo_Dynamic

/**
* 定义dp[i][j] = 机器人走到终点的路径数
* 初始值dp[0][*] = 1
* 初始值dp[*][0] = 1
* 初始值dp[1][1] = dp[0][1]+dp[1][0]
* 初始值dp[2][1] = dp[][1]+dp[2][0]
* 推出dp[n]的公式：
* dp[i][j] = dp[i-1][j]+dp[i][j-1]
 */

func robotGo(c, r int) int {
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
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[c-1][r-1]
}
