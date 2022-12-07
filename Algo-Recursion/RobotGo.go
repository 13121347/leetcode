package Algo_Recursion

func robotGoReverse(m, n int) int {
	res := make([][]int, m)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if res[i][j] != 0 {
			return res[i][j]
		}
		if i == 0 || j == 0 {
			return 1
		}
		res[i][j] = dfs(i-1, j) + dfs(i, j-1)
		return res[i][j]
	}
	return dfs(m, n)
}
