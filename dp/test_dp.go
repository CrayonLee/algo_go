package dp

//最长递增子序列

//暴力解法
//func lengthOfLIS(nums []int) int {
//	dp := make([]int, len(nums))
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//	}
//	res := 0
//	for _, v := range dp {
//		res = max(res, v)
//	}
//	return res
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//二分解法

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))
	piles := 0

	for _, v := range nums {
		pocker := v
		left, right := 0, piles
		for left < right {
			mid := (left + right) >> 1
			if top[mid] >= pocker {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = pocker
	}
	return piles
}
