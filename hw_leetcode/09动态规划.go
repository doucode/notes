package hw_leetcode

//9. 动态规划 总结：https://mp.weixin.qq.com/s/Cw39C9MY9Wr2JlcvBQZMcA
//代表题目：213🉑 触类旁通：1043🉑.416🉑.123🉑.62🉑

//	62. 不同路径	https://leetcode.cn/problems/unique-paths/
//一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//输入：m = 3, n = 7
//输出：28
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

//
func uniquePaths(m int, n int) int {

}

//

//	213. 打家劫舍 II		https://leetcode.cn/problems/house-robber-ii/
//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//输入：nums = [2,3,2]
//输出：3
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(rb(nums[:n-1]), rb(nums[1:])) // 在[0,n-2] 和[1,n-1]两个区间里找
}

func rb(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second) //dp[i]=max(dp[i−2]+nums[i],dp[i−1])
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	416. 分割等和子集	https://leetcode.cn/problems/partition-equal-subset-sum/
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11]
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	k := sum >> 1
	dp := make([]bool, k+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := k; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[k]
}

//	123. 买卖股票的最佳时机 III	https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//输入：prices = [3,3,5,0,0,3,1,4]
//输出：6
func maxProfit(prices []int) int {
	buy1, sell1 := prices[0], 0
	buy2, sell2 := prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = min(buy1, prices[i])
		sell1 = max(sell1, prices[i]-buy1)
		buy2 = min(buy2, prices[i]-sell1) // 第二次购买的价格 = 股票价格 - 第一次的收益
		sell2 = max(sell2, prices[i]-buy2)
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//	1043. 分隔数组以得到最大和	https://leetcode.cn/problems/partition-array-for-maximum-sum/
//输入：arr = [1,15,7,9,2,5,10], k = 3
//输出：84
//解释：
//因为 k=3 可以分隔成 [1,15,7] [9] [2,5,10]，结果为 [15,15,15,9,10,10,10]，和为 84，是该数组所有分隔变换后元素总和最大的。
//若是分隔成 [1] [15,7,9] [2,5,10]，结果就是 [1, 15, 15, 15, 10, 10, 10] 但这种分隔方式的元素总和（76）小于上一种。
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		curMax := 0
		for j := 1; j <= min(i, k); j++ {
			curMax = max(curMax, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+curMax*j)
		}
	}
	return dp[n]
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
