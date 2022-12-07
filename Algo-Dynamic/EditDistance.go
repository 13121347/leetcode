package Algo_Dynamic

func editorDis(word1, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)

	word1Arr := []rune(word1)
	word2Arr := []rune(word2)

	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	dp[0][0] = 0
	//dp[0][0...n2]的初始值
	for j := 1; j < n2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	//dp[0...n2][0]的初始值
	for i := 1; i < n1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	//公式推出dp[n1][n2]
	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			if word1Arr[i-1] == word2Arr[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = triMin(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[n1][n2]
}

func triMin(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
