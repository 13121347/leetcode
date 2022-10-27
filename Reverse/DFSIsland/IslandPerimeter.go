package DFSIsland

//岛屿周长问题：https://leetcode.cn/problems/island-perimeter/

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) int {
	ans := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if !inArea(grid, r, c) {
			ans++
			return
		}

		if grid[r][c] == 0 {
			ans++
			return
		}

		if grid[r][c] != 1 {
			return
		}
		grid[r][c] = 2

		for _, d := range dir4 {
			dfs(r+d.x, c+d.y)
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				dfs(r, c)
			}
		}
	}
	return ans
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
