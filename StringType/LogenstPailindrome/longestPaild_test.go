package LogenstPailindrome

import (
	"fmt"
	"testing"
	"time"
)

func Test_longestPalindrome(t *testing.T) {
	startedAt := time.Now()

	time.Sleep(time.Second)
	defer func(startedAt time.Time) { fmt.Println("3:", time.Since(startedAt)) }(startedAt)
	//defer fmt.Println("2:",time.Since(startedAt))
	//defer func() { fmt.Println("1:",time.Since(startedAt)) }()
	time.Sleep(time.Second)
}
