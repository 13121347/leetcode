package SliceWindow

//存在重复元素 https://leetcode.cn/problems/contains-duplicate-ii/
//easy

func containsNearbyDuplicate(nums []int, k int) bool {
	//滑动窗口两个值，扩大窗口最多到size = K
	//用map记录窗口中的数字
	start, end := 0, 0
	existMap := make(map[int]struct{})
	for end < len(nums) {
		_, ok := existMap[nums[end]]
		if ok {
			return true
		} else {
			existMap[nums[end]] = struct{}{}
		}
		if end-start+1 > k {
			delete(existMap, nums[start])
			start++
		}
		end++
	}
	return false
}
