package SpiralOrder

//螺旋矩阵https://leetcode.cn/problems/spiral-matrix/description/
//https://leetcode.cn/problems/spiral-matrix-ii/solutions/12594/spiral-matrix-ii-mo-ni-fa-she-ding-bian-jie-qing-x/

func spiralOrder1(matrix [][]int) (res []int) {
	//判断边界条件
	if len(matrix) == 0 {
		return
	}

	//计算matrix中的数据总数，用于输出边界控制
	lengthR := len(matrix)
	lengthC := len(matrix[0])
	numCount := lengthC * lengthR

	//定义遍历层
	leftMargine, rightMargin, topMargine, belowMargine := 0, lengthC-1, 0, lengthR-1

	//开始遍历
	for numCount > 0 {
		//从左到右
		for i := leftMargine; i <= rightMargin && numCount > 0; i++ {
			res = append(res, matrix[topMargine][i])
			numCount--
		}
		//向下收缩
		topMargine++

		//从上到下
		for i := topMargine; i <= belowMargine && numCount > 0; i++ {
			res = append(res, matrix[i][rightMargin])
			numCount--
		}
		//向左收缩
		rightMargin--

		//从右到左
		for i := rightMargin; i >= leftMargine && numCount > 0; i-- {
			res = append(res, matrix[belowMargine][i])
			numCount--
		}
		//向上收缩
		belowMargine--

		//从下到上
		for i := belowMargine; i >= topMargine && numCount > 0; i-- {
			res = append(res, matrix[i][leftMargine])
			numCount--
		}
		//向右收缩
		leftMargine++
	}
	return
}
