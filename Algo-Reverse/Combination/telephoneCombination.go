package Combination

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}

	var path []byte //这里注意byte和string的替换
	var backTrack func(index int)
	backTrack = func(index int) {
		//终止条件
		if len(path) == len(digits) {
			comStr := string(path)
			ans = append(ans, comStr)
			return
		}
		//获取映射字母
		letter := phoneMap[string(digits[index])]
		//遍历
		for i := 0; i < len(letter); i++ {
			path = append(path, letter[i])
			backTrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backTrack(0)
	return
}
