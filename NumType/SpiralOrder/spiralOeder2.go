package SpiralOrder

//生成螺旋矩阵 https://leetcode.cn/problems/spiral-matrix-ii/
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素

func generateMatrix(n int) [][]int {
	leftMargine, rightMargine, topMargine, belowMargint := 0, 0, 0, 0
	newMartrix := make([][]int, n)
	for i := 0; i < n; i++ {
		newMartrix[i] = make([]int, n)
	}

	num, target := 1, n*n

	for num <= target {
		//从左到右
		for i := leftMargine; i <= rightMargine; i++ {
			newMartrix[topMargine][i] = num
			num++
		}
		topMargine++
		//从上到下
		for i := topMargine; i < belowMargint; i++ {
			newMartrix[rightMargine][i] = num
			num++
		}
		rightMargine--
		//从右到左
		for i := rightMargine; i >= leftMargine; i++ {
			newMartrix[belowMargint][i] = num
			num++
		}
		belowMargint--
		//从下到上
		for i := belowMargint; i >= topMargine; i++ {
			newMartrix[i][leftMargine] = num
			num++
		}
		leftMargine++
	}
	return newMartrix
}
