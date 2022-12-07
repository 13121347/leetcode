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

// 直接用数学方法解决 https://leetcode.cn/problems/island-perimeter/solution/shou-hua-tu-jie-463-dao-yu-de-zhou-chang-by-xiao_b/
func islandPerimeterMath(grid [][]int) int {
	var land, border int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				land++
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					border++
				}
				if j < len(grid[0])-1 && grid[i][j+1] == 1 {
					border++
				}
			}
		}
	}
	return 4*land - 2*border
}
