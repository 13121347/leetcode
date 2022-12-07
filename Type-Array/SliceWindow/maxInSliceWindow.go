package SliceWindow

import (
	"container/heap"
	"sort"
)

//滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/

var a []int

type hp struct {
	sort.IntSlice
}

//重新定义的这个Less函数，虽然堆里放的元素是index,但比较的时候用的是num
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	//初始化窗口大的堆
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		//这里是把下标放到heap中
		q.IntSlice[i] = i
	}
	heap.Init(q)

	//定义返回结果
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]

	//窗口开始滑动，i:右边框
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
