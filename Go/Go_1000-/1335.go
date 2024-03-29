// Go
// leetcode 1335
// https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/

func minDifficulty(a []int, d int) int {
	n := len(a)
	if n < d {
		return -1
	}

	dp := make([][]int, d)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		p := &dp[i][j]
		if *p != -1 { // 之前计算过了
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if i == 0 {                 // 只有一天，必须完成所有工作
			for _, x := range a[:j+1] {
				res = max(res, x)
			}
			return
		}
		res = math.MaxInt
		mx := 0
		for k := j; k >= i; k-- {
			mx = max(mx, a[k]) // 从 a[k] 到 a[j] 的最大值
			res = min(res, dfs(i-1, k-1)+mx)
		}
		return
	}
	return dfs(d-1, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}