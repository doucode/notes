package hw_leetcode

//单调栈 总结：单调栈的性质，看人家总结。https://mp.weixin.qq.com/s/YeQ7eE0-hZpxJfJJziq25Q
//代表题目：85🆗  触类旁通：739🉑.503🉑.84🉑.42🉑.901🉑.239🉑

// 42.接雨水  https://leetcode.cn/problems/trapping-rain-water/
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
func trap(height []int) int {
	stack := []int{}
	ans := 0

	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curHeight * curWidth
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//	239. 滑动窗口最大值	https://leetcode.cn/problems/sliding-window-maximum/
//	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//	你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//	输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//	输出：[3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{} //	单调递减队列，q[0]为窗口内最大值
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

//	739. 每日温度	https://leetcode.cn/problems/daily-temperatures/
//	给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，
//	其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。
//	如果气温在这之后都不会升高，请在该位置用0 来代替。
//	输入: temperatures = [73,74,75,71,69,72,76,73]
//	输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(temperatures []int) []int {
	//单调栈
	ans := make([]int, len(temperatures))
	var stack []int // 单调递减栈
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			idx := stack[len(stack)-1]
			ans[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 	84. 柱状图中最大的矩形	https://leetcode.cn/problems/largest-rectangle-in-histogram/
// 	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 	求在该柱状图中，能够勾勒出来的矩形的最大面积。
//	输入：heights = [2,1,5,6,2,3]
//	输出：10
func largestRectangleArea(heights []int) int {
	ans := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || i == n-1 {
			return 0
		}
		return heights[i-1]
	}

	stack := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && getHeight(stack[len(stack)-1]) > getHeight(i) {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop stack
			ans = max(ans, getHeight(idx)*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i) // push stack
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//	503. 下一个更大元素 II	https://leetcode.cn/problems/next-greater-element-ii/
//	给定一个循环数组nums（nums[nums.length - 1]的下一个元素是nums[0]），返回nums中每个元素的 下一个更大元素 。
//
//	数字 x的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
//	如果不存在，则输出 -1
//
//	输入: nums = [1,2,1]
//	输出: [2,-1,2]
func nextGreaterElements(nums []int) []int {
	ans := make([]int, 0)
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		ans = append(ans, -1)
	}
	for i := 0; i < len(nums)*2; i++ {
		num := nums[i%len(nums)]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			index := stack[len(stack)-1]
			ans[index] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}
	return ans
}

//	85. 最大矩形		https://leetcode.cn/problems/maximal-rectangle/
//	给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//	输入：matrix = [["1","0","1","0","0"],
//				   ["1","0","1","1","1"],
//				   ["1","1","1","1","1"],
//				   ["1","0","0","1","0"]]
//输出：6
func maximalRectangle(matrix [][]byte) (ans int) {

}

//	901. 股票价格跨度		https://leetcode.cn/problems/online-stock-span/
//	编写一个 StockSpanner 类，它收集某些股票的每日报价，并返回该股票当日价格的跨度。
//	今天股票价格的跨度被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
//
//	例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是 [1, 1, 1, 2, 1, 4, 6]。
type StockSpanner struct {
	que, stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{[]int{}, []int{}}
}

func (this *StockSpanner) Next(price int) int {
	ans := 0
	this.que = append(this.que, price)
	idx := len(this.que) - 1
	for len(this.stack) > 0 && this.que[this.stack[len(this.stack)-1]] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}

	if len(this.stack) == 0 {
		this.stack = append(this.stack, idx)
		return idx + 1
	}

	ans = idx - this.stack[len(this.stack)-1]
	this.stack = append(this.stack, idx)
	return ans
}
