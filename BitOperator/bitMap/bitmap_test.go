package bitMap

import (
	"fmt"
	"testing"
)

func Test_bitMap(t *testing.T) {
	mybitMap := InitBitMap(3)
	mybitMap.add(1)
	mybitMap.add(2)
	mybitMap.add(3)
	mybitMap.add(4)

	fmt.Println(mybitMap.contain(1))
	fmt.Println(mybitMap.contain(3))
	fmt.Println(mybitMap.contain(4))
	fmt.Println(mybitMap.contain(5))
	fmt.Println(mybitMap.contain(7))
}
