package Partition

import "fmt"

//分割回文串 https://leetcode.cn/problems/palindrome-partitioning/

func partition(s string) (ans [][]string) {
	var path []string

	var backTrack func(index int)
	backTrack = func(index int) {
		//终止条件
		if index == len(s) {
			ans = append(ans, append([]string(nil), path...))
			fmt.Println("ans=", ans)
			return
		}
		//循环
		for i := index; i < len(s); i++ {
			fmt.Println("index=", index, "i=", i, " path = ", path)

			if isPalidrome(s, index, i) {
				path = append(path, s[index:i+1])
				fmt.Println(" path = ", path)
				backTrack(i + 1) //这里为什么是i+1,不是index+1?----这样可以一直往后面找数，而不是往前
				path = path[:len(path)-1]
			}
		}
	}

	backTrack(0)
	return
}

func isPalidrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true

}
