package Combination

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
// https://leetcode.cn/problems/combinations/
// medium

import "fmt"

func combinationNK(n, k int) (ans [][]int) {
	//符合条件的结果
	path := []int{}
	//markedArr := make([]bool,n+1)
	var backStack func(index int)
	backStack = func(index int) { //这里的index相当于横向递归
		//结束条件
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		//递归
		for i := index; i <= n; i++ { //这里相当于纵向递归,每次从index往后，就是一层一层处理
			fmt.Println("index=", index, "i=", i)
			//if markedArr[i] == false {
			path = append(path, i)
			//markedArr[i] = true
			backStack(i + 1) //这里使用i+1,而不是index+1，只往后面找数据，所以这里不需要有marked记录。原因是题目中【3，4】 = 【4，3】即组合问题，不是全排列问题
			path = path[:len(path)-1]
			//markedArr[i] = false
			//}
		}
	}
	backStack(1)
	return
}

func combinationNKCut(n, k int) (ans [][]int) {
	//符合条件的结果
	path := []int{}
	//	markedArr := make([]bool,n+1)
	var backStack func(index int)
	backStack = func(index int) { //这里的index相当于横向递归
		//结束条件
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		//递归
		for i := index; i <= n-(k-len(path))+1; i++ { //这里相当于纵向递归,每次从index往后，就是一层一层处理。在这里剪枝
			//if markedArr[i] == false {
			path = append(path, i)
			//markedArr[i] = true
			backStack(i + 1) //这里使用i+1,而不是index+1，只往后面找数据，所以这里不需要有marked记录。原因是题目中【3，4】 = 【4，3】
			path = path[:len(path)-1]
			//markedArr[i] = false
			//}
		}
	}
	backStack(1)
	return
}
