package Dynamic

import "fmt"

//通配符匹配，. *  https://leetcode.cn/problems/regular-expression-matching/
//https://leetcode.cn/problems/regular-expression-matching/solution/by-lebron-s-u27g/
func isMatched2(s, p string) bool {
	//定义dp,dp[i][j]表示，s前i个字符，p前j个字符匹配的结果
	lenS, lenP := len(s), len(p)
	dp := make([][]bool, lenS+1)
	//构造dp
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]bool, lenP+1)
	}

	//初始化dp
	//s,p都为0时，= true
	//sp不同时为0时=false，默认值，不用修改
	dp[0][0] = true
	for i := 1; i <= lenP; i++ {
		if p[i-1] == '*' && i >= 2 {
			dp[0][i] = dp[0][i-2]
		}

	}

	//状态转移方程：
	/**
	s，p都是字母时，dp[i][j] = dp[i-1][j-q] && s[i-1]==p[j-1]
	p 为 '.'时，dp[i][j] = dp[i-1][j-1]
	p 为 '*'时，dp[i][j] = dp[i-1]dp[j-1]&&s[i]==p*前面的字母？
	*/
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if p[j-1] == '.' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
				fmt.Println("dp[", i, "]", "[", j, "]=", dp[i][j])
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				fmt.Println("dp[", i, "]", "[", j, "]=", dp[i][j])
				if p[j-1] == s[i] {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
					fmt.Println("dp[", i, "]", "[", j, "]=", dp[i][j])
				}
			}
		}
	}

	return dp[lenS][lenP]
}

func isMatchRight(s string, p string) bool {
	m, n := len(s), len(p)
	//辅助函数，s和p的第i,j个字符是否可以匹配。s[i-1]==p[j-1]或者p[j-1]='.'
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	//构建动态数组
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}

	//初始条件 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
	f[0][0] = true

	//动态转移方程

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
