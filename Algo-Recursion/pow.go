package Algo_Recursion

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	return pow(x, n-1) * x
}

//快速幂+递归
func mypow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	} else {
		return 1 / quickPow(x, n)
	}
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x

}
