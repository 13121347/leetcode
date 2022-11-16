package DoublePointer

// 移动0 https://leetcode.cn/problems/move-zeroes/

//移到前面
func moveZeroesFront(nums []int) {
	//定义before，after,before用来遍历数组，指到0时进行处理。
	//after指向数组前段为0的数据前一个数据
	//处理时将nums[after]和为0的nums[before]进行数据交换，同时after移动，否则after不移动
	after, before := 0, 0

	for ; before < len(nums)-1; before++ {
		for nums[after] == 0 {
			after++
		}
		if before < after {
			before = after
		}
		if nums[before] == 0 {
			nums[after], nums[before] = nums[before], nums[after]
			after++
		}
	}
}

//左右指针的交换是不保序的，要保证原顺序，用前后指针
//移到后面,交换，不保证顺序
func moveZeroesBack(nums []int) {
	//定义左右两个指针进行遍历
	//left处理遍历，找到数字0进行处理
	//right指向非0，用于交换，right右边数值全为0
	left, right := 0, len(nums)-1

	for left < right {
		for nums[right] == 0 {
			right--
		}

		if nums[left] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
		left++
	}
}

//移到后面,保证顺序
func moveZeroesBackOrder(nums []int) {
	front, after := 0, 0

	for ; front < len(nums); front++ {
		if nums[front] != 0 {
			nums[after] = nums[front]
			after++
		}
	}
	//将末尾元素置为0
	for ; after < len(nums); after++ {
		nums[after] = 0
	}
}
