/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package leetcode

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2

	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

// @lc code=end
