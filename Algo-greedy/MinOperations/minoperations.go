package MinOperations

// https://leetcode.cn/problems/minimum-operations-to-make-array-equal/description/
// 给你一个整数 n，即数组的长度。请你返回使数组 arr 中所有元素相等所需的 最小操作数（arr[i] = (2 * i) + 1，
// 一次操作中，可以选出两个下标，记作 x 和 y ，并使 arr[x] 减去 1 、arr[y] 加上 1）
// medium

//分析：本题给定的操作并不会使数组中所有元素之和变化，且题目要求让所有元素相等，那么数组中所有元素的平均值即为最后数组中每一个元素的值

//根据题目 arr[i] = (2 * i) + 1，这些数加起来除n得到的平均数 = n (是有这么个数学公式的)，且操作是一加一减算一次
func minOperations(n int) (ans int) {
	for i := 0; i < n; i++ {
		x := 2*i + 1
		if x > n {
			ans = ans + x - n
		}
	}
	return
}
