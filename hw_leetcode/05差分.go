package hw_leetcode

//5. 差分   总结：https://mp.weixin.qq.com/s/123QujqVn3--gyeZRhxR-A
//代表题目：1094🉑  触类旁通：1109🉑.121🉑.122🉑.253🔒

//	121. 买卖股票的最佳时机		https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
//给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//输入：[7,1,5,3,6,4]
//输出：5
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	max, minPrice := -1<<31, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > max { // prices[i] - minPrice为第i日卖出的最大值
			max = prices[i] - minPrice
		}
	}
	return max
}

//	122. 买卖股票的最佳时机 II		https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
//给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。
//
//在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
//返回 你能获得的 最大 利润。
//输入：prices = [7,1,5,3,6,4]
//输出：7
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1]) //x个长度为1的区间的 正差值 累加和
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	1109. 航班预订统计	https://leetcode.cn/problems/corporate-flight-bookings/
//输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//输出：[10,55,45,25,25]
//解释：
//航班编号        1   2   3   4   5
//预订记录 1 ：   10  10
//预订记录 2 ：       20  20
//预订记录 3 ：       25  25  25  25
//总座位数：      10  55  45  25  25
//因此，answer = [10,55,45,25,25]
func corpFlightBookings(bookings [][]int, n int) []int {
	// 差分
	ans := make([]int, n)
	for _, booking := range bookings {
		ans[booking[0]-1] += booking[2]
		if booking[1] < n {
			ans[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

//	1094. 拼车	https://leetcode.cn/problems/car-pooling/
//车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）
//
//给定整数capacity和一个数组 trips , trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi 乘客，接他们和放他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。
//
//当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false。
//输入：trips = [[2,1,5],[3,3,7]], capacity = 4
//输出：false
func carPooling(trips [][]int, capacity int) bool {
	psgs := make([]int, 1001)
	for _, t := range trips {
		psgs[t[1]] += t[0]
		psgs[t[2]] -= t[0]
	}

	total := 0
	for _, psg := range psgs {
		total += psg
		if total > capacity {
			return false
		}
	}
	return true
}
