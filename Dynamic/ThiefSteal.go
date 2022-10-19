package Dynamic

func thief(houses []int) int {
	if len(houses) == 0 {
		return 0
	}
	dp := make([]int, len(houses))
	dp[0] = houses[0]
	dp[1] = max(dp[0], houses[1])

	for i := 2; i < len(houses); i++ {
		dp[i] = max(dp[i-2]+houses[i], dp[i-1])
	}
	return dp[len(houses)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
