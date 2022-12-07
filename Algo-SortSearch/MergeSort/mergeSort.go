package MergeSort

//动画讲解 http://wjhsh.net/wangyiyang-p-12205574.html
//https://blog.csdn.net/weixin_46040684/article/details/121547182
//时间复杂度非常稳定在任何情况都是O(nlogn)
func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	//从中间拆分
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	//合并
	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	//定义左边部分开始遍历的index
	leftIndex := start
	//定义右边部分开始遍历的index
	rightIndex := mid + 1
	//定义一个临时数组
	tmpIndex := 0
	tmp := make([]int, 1+end-start)

	//遍历挑选
	for leftIndex <= mid && rightIndex <= end {
		if nums[leftIndex] <= nums[rightIndex] {
			tmp[tmpIndex] = nums[leftIndex]
			tmpIndex++
			leftIndex++
		} else {
			tmp[tmpIndex] = nums[rightIndex]
			tmpIndex++
			rightIndex++
		}
	}

	//判断哪一边的序列还有剩余元素
	var appendStart, appendEnd int
	//left部分已经处理完
	if leftIndex > mid {
		appendStart = rightIndex
		appendEnd = end
	} else { //right部分处理完
		appendStart = leftIndex
		appendEnd = mid
	}

	//将剩余元素放入tmp
	for appendStart <= appendEnd {
		tmp[tmpIndex] = nums[appendStart]
		tmpIndex++
		appendStart++
	}

	//将tmp放入原数组
	for k, v := range tmp {
		nums[start+k] = v
	}
}
