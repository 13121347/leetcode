package DFSIsland

func numIslands(grid [][]byte) int {
	//定义返回结果
	num := 0
	//定义dfs处理函数
	var dfsNumber func(r, c int)
	dfsNumber = func(r, c int) {
		//格子越界了，返回
		if !inAreaB(grid, r, c) {
			return
		}

		//不是陆地，返回
		if grid[r][c] != '1' {
			return
		}

		//遍历过的格子修改数值
		grid[r][c] = '2'

		//dfs周边格子
		for _, v := range dir {
			dfsNumber(r+v.x, c+v.y)
		}
	}

	//循环处理整个数组
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				dfsNumber(r, c)
				num++
			}
		}
		//每次跳出来dfs相当于一个岛屿结束

	}
	return num
}

func inAreaB(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
