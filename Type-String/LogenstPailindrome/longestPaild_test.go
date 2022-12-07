package LogenstPailindrome

import (
	"fmt"
	"testing"
)

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i //向channel中写入数据
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in //从in channel中读数据，过滤后再写入out channel
		//if i%prime != 0{
		out <- i
		//}
	}
}

func Test_longestPalindrome(t *testing.T) {
	/*	startedAt := time.Now()

		time.Sleep(time.Second)
		defer func(startedAt time.Time) { fmt.Println("3:", time.Since(startedAt)) }(startedAt)
		//defer fmt.Println("2:",time.Since(startedAt))
		//defer func() { fmt.Println("1:",time.Since(startedAt)) }()
		time.Sleep(time.Second)*/
	defer A()
	defer B()
	defer C()
	panic("panic main")
	fmt.Println("func main")
}

func A() {
	x := recover()
	fmt.Println("func A,recover:", x)
	panic("panic A")
}

func B() {
	panic("panic B")
	fmt.Println("func B")
}

func C() {
	fmt.Println("func C")
}
