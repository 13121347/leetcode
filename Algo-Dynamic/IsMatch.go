package Algo_Dynamic

//通配符匹配 https://leetcode.cn/problems/wildcard-matching/
func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	//dp[i][j]表示，s前i个字符和p前j个字符的匹配情况
	dp := make([][]bool, lenP+1)
	for i := 0; i <= lenP; i++ {
		dp[i] = make([]bool, lenS+1)
	}
	//初始状态，
	//dp[0][0]:sp都为空时，=true
	//dp[i][0]:p为空时，无法匹配，=false
	//dp[0][j]:p[j]='*'时才=true
	//区分dp下标和p,s，下标的不同
	dp[0][0] = true
	for i := 0; i <= lenS; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[lenS][lenP]
}
