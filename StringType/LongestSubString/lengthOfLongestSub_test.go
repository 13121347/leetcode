package LongestSubString

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "addbfhjiauufiuejs"
	result1 := lengthOfLongestSubstring(s)
	fmt.Println(result1)

	result2 := lengthOfLongestSubstring2(s)
	fmt.Println(result2)
}
