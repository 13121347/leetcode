package DFSIsland

import (
	"math"
)

//岛屿的最大面积 https://leetcode.cn/problems/max-area-of-island/

type pira struct {
	x int
	y int
}

var dir = []pira{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	maxArea := 0
	var getArea func(r, c int)
	getArea = func(r, c int) {
		if !inArea(grid, r, c) {
			//res=0 第一次这么写，每次都把res重置为0了
			return
		}
		if grid[r][c] != 1 {
			//res=0
			return
		}

		if grid[r][c] == 1 {
			res++
		}

		grid[r][c] = 2

		for _, v := range dir {
			getArea(r+v.x, c+v.y)
		}

	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				getArea(r, c) //函数如果从这里退出，说明一块连着的岛已经遍历结束
				maxArea = int(math.Max(float64(maxArea), float64(res)))
			}
			res = 0 //第二次没有加这一句，导致计算出的是所有陆地面积了，因为闭包函数里，res是一直在递增的
		}
	}
	return maxArea
}
