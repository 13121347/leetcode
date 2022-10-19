package Dynamic

/**
*  需要凑X元，可选的钱从大到小排序是a,b,c，这里有三个概念，X元，abc种类，硬币数n
*  定义dp[n] = 兑换X元时的最少的硬币数
*  dp[n-1] = 兑换X-（a/b/c）最少的硬币数
* 所以 dp[n]=dp[n-1]+(a/b/c)
* 但是
 */
