package tree

func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for left := 0; left < i; left++ {
			right := i - 1 - left
			dp[i] += dp[left] * dp[right]
		}
	}

	return dp[n]
}
