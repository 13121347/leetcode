package MyStack

//有效的括号 https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	matchedMap := make(map[byte]byte)
	matchedMap['}'] = '{'
	matchedMap[')'] = '('
	matchedMap[']'] = '['

	brackToMatched := make([]byte, len(s))
	for i, _ := range s {
		if _, ok := matchedMap[s[i]]; ok {
			if len(brackToMatched) == 0 || brackToMatched[len(brackToMatched)-1] != matchedMap[s[i]] {
				return false
			}
			brackToMatched = brackToMatched[:len(brackToMatched)-1]
		} else {
			brackToMatched = append(brackToMatched, s[i])
		}
	}
	return len(brackToMatched) == 0
}
