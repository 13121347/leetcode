package LogenstPailindrome

//最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/description/
//中心扩展法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	//遍历字符串上每一个字母，以它为中心，向两边扩展，遇到不相同则返回
	for i := 0; i < len(s); i++ {
		//奇数中心
		left1, right1 := expandAroundCenter(s, i, i)
		//偶数中心
		left2, right2 := expandAroundCenter(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left = left - 1
		right = right + 1
	}
	return left + 1, right - 1
}
