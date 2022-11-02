package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSortDown(t *testing.T) {
	testArr := []int{5, 2, 3, 8, 4, 9}
	QuickSortUp(testArr, 0, 5)
	for _, k := range testArr {
		fmt.Print(k, " ")
	}
	fmt.Println()
	fmt.Println("***************")
	QuickSortDown(testArr, 0, 5)
	for _, k := range testArr {
		fmt.Print(k, " ")
	}
}

func TestSortUsage(t *testing.T) {
	m1 := MySort{
		Id: 1,
		s:  1,
	}

	m2 := MySort{
		Id: 3,
		s:  1,
	}

	m3 := MySort{
		Id: 5,
		s:  1,
	}

	msarr := []MySort{m1, m2, m3}

	for _, v := range msarr {
		fmt.Print(v, "  ")
	}
}
