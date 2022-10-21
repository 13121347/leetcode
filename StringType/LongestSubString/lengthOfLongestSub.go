package LongestSubString

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	//转到数组中
	//记录两个下标，对下标中间的字符使用map判断是否存在，以及存在位置
	//当移动后面的下标后会引入重复字符，进行缩短

	cMap := map[byte]int{}
	n := len(s)

	right, ans := -1, 0
	//依次遍历s
	for i := 0; i < n; i++ {
		//删除map里之前存在的，遍历过的字符
		if i != 0 {
			delete(cMap, s[i-1])
		}
		//左指针没有指向最后一个字符 && map里面没有出现过这个字符，s[right+1] = 当前的字符串，是map的key
		for right+1 < n && cMap[s[right+1]] == 0 {
			//在map里记录这个字符
			cMap[s[right+1]]++
			right++
		}
		//到头了或者出现重复的字符了，看看当前[指针位置，遍历起点]长度
		ans = max(ans, right-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring2(s string) int {
	movedIndex := 0
	markedMap := map[byte]int{}
	ans := 0
	for sIndex, _ := range s {
		if sIndex != 0 {
			delete(markedMap, s[sIndex-1])
		}

		for movedIndex < len(s) && markedMap[s[movedIndex]] == 0 {
			markedMap[s[movedIndex]] = 1
			movedIndex++
		}

		ans = max(ans, movedIndex-sIndex)
	}

	return ans
}
