package RemoveKNumber

// 移掉K位数字 https://leetcode.cn/problems/remove-k-digits/
// medium 贪心+单调栈

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)

	}
	return ""
}
