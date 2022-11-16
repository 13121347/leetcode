package SliceWindow

import "math"

//长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/
//medium
/*
	这道题是一道经典的滑动窗口问题[4]。
	使用start、end指针，分别表示滑动窗口的起始、终止位置
	移动end指针，扩大窗口，直到子数组达到目标值target
	移动start指针，缩小窗口，直到子数组不再满足>=target
*/

func minSubArrayLen(target int, nums []int) int {
	//增加异常处理
	if len(nums) == 0 {
		return 0
	}
	//maxlength初始值为math.MaxInt32
	start, end, maxLenngth := 0, 0, math.MaxInt32
	sum := 0
	//我原本的写法相当于把嵌套的for循环用两个if替换，通过外面的循环一直处理
	for end < len(nums) {
		//sum < target,扩大窗口
		/*if sum < target{
			//扩大窗口
			end++
		}*/
		//sum >=0 记录当前窗口长度，缩小窗口 这里之前写的是if，逻辑应该是如果sum>target，就要一直收缩窗口
		sum = sum + nums[end]
		for sum >= target {
			currentWindow := end - start + 1
			maxLenngth = min(maxLenngth, currentWindow)
			sum = sum - nums[start]
			start++
		}
		//这里也不用判断sum<target才++，因为上面改成了for，从for出来之后，必然sum<target
		end++
	}
	//处理没有结果的情况
	if maxLenngth == math.MaxInt32 {
		return 0
	}
	return maxLenngth
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//[1,4,4]
func minSubArrayLenSelf(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end, sum := 0, 0, 0

	for end < len(nums) {
		if sum < target {
			sum = sum + nums[end]
			//扩大窗口
			end++ //end在这里++，会影响后面缩小窗口时候current的计算
		}
		//sum >=0 记录当前窗口长度，缩小窗口
		if sum >= target {
			currentWindow := end - start + 1
			ans = min(ans, currentWindow)
			sum = sum - nums[start]
			start++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
