package Partition

import (
	"strings"
)

func restoreIpAddresses(s string) (ans []string) {

	if len(s) > 12 || len(s) < 4 {
		return
	}
	var path []string
	//剩余段数
	residue := 4

	var backTrack func(index, residue int)

	backTrack = func(index, residue int) {
		if index == len(s) && residue == 0 {
			ip := strings.Join(path, ".")
			ans = append(ans, ip)
			return
		}

		//每段取三个
		for i := index; i < index+3; i++ {
			if i >= len(s) {
				break
			}
			//减枝
			if residue*3 < len(s)-i {
				continue
			}

			if isIpSegment(s, index, i) {
				currentIpSegment := s[index : i+1]
				path = append(path, currentIpSegment)
				backTrack(i+1, residue-1)
				path = path[:len(path)-1]
			}
		}
	}

	backTrack(0, residue)
	return
}

//判断字串是否符合ip要求
func isIpSegment(s string, left, right int) bool {
	//首位0情况
	if right > left && s[left] == '0' {
		return false
	}
	//判断对应数字是否满足范围
	num := 0
	for i := left; i <= right; i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num >= 0 && num <= 255
}
