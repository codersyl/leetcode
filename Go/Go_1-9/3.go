// Go
// leetcode 3
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	vis, n, res := make([]bool, 128), len(s), 0
	for i, j := 0, 0; j < n; j++ {
		idx_j := int(s[j])
		if vis[idx_j] {
			for int(s[i]) != idx_j {
				vis[int(s[i])] = false
				i++
			}
			i++
		} else {
			vis[idx_j] = true
		}
		res = max(j-i+1, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}