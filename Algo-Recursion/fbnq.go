package Algo_Recursion

func fbnc(index int) int {
	if index <= 1 {
		return 1
	}
	return fbnc(index-1) + fbnc(index-2)
}
