package Dynamic

//https://zhuanlan.zhihu.com/p/91582909

/**
* 定义dp[i] = 跳上n级台阶的方法
* 初始值dp[1] = 1,跳上第1个台阶有1个方法：一次跳1个台阶，跳1次
* 初始值dp[2] = 2,跳上第2个台阶有2个方法：一次跳1个台阶，跳2次  + 一次跳2个台阶，跳1次
* 推出dp[n]的公式：
* 如果蛙蛙想到第3个台阶，那么他可以先跳到第1个台阶，然后dp[2]的方式到第三个台阶，或者先到第2个台阶，然后dp[1]的方式到第三个台阶
* 就有公式dp[3] = dp[2]+dp[1]
* 如果蛙蛙想到第4个台阶，那么他可以先跳到第1个台阶，然后dp[3]的方式到第三个台阶，或者先到第2个台阶，然后dp[2]的方式到第三个台阶
* 就有公式dp[4] = dp[3]+dp[2]
* 一次类推，dp[n] = dp[n-1]+dp[n-2]
 */

func frogJump(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		//注意这里的取模操作
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
