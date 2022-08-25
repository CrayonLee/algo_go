package dp

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word1)

	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)

	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= 1; i++ {
		dp[i] = make([]int, l1+1)
		dp[i][0] = i

		for j := 1; j <= l2; j++ {

			if word1[i-1:i] == word2[j-1:j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minThree(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[l1][l2]
}

func minThree(a, b, c int) int {
	min := 0
	if a < b {
		min = a
	} else {
		min = b
	}

	if min > c {
		min = c
	}
	return min
}
