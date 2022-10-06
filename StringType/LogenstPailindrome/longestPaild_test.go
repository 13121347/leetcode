package LogenstPailindrome

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	testStr := "cbbc"
	result := longestPalindrome(testStr)
	fmt.Println(result)
}
