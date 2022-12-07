package NQueen

// N皇后问题 https://leetcode.cn/problems/n-queens/
func solveNQueens(n int) [][]string {
	return nil
}

//判断当前位置是否可以落子
func isValid(board [][]int, n, row, col int) bool {
	//行不用判断，每层只有一个
	//列判断
	for c := 0; c < n; c++ {
		if board[c][col] == 'Q' {
			return false
		}
	}
	//主对角线判断（45度）
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	//副对角线（135度）
	i, j = row-1, col+1
	for i >= 0 && j <= n-1 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}

//转换结果字符串列表
func transBoard(board [][]int) (path []int) {
	for col := 0; col < len(board); col++ {
		for row := 0; row < len(board[0]); row++ {
			//c := strconv.Itoa((board[row][row]))
		}
	}
	return nil
}
