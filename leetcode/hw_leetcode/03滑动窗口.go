package hw_leetcode

// 滑动窗口   总结：https://mp.weixin.qq.com/s/ioKXTMZufDECBUwRRp3zaA
//代表题目：1208🉑.3🉑.1004🉑   触类旁通：340🔒.1151🔒.159🔒.1100🔒

//	3. 无重复字符的最长子串	https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	l, r, ans := 0, 0, 0
	window := make(map[byte]int, len(s))
	for r < len(s) {
		if idx, ok := window[s[r]]; ok && idx >= l { // 判断左边界
			l = idx + 1
		}
		window[s[r]] = r
		r++
		ans = max(ans, r-l)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	1208. 尽可能使字符串相等	https://leetcode.cn/problems/get-equal-substrings-within-budget/
//	输入：s = "abcd", t = "bcdf", maxCost = 3
//	输出：3
func equalSubstring(s string, t string, maxCost int) int {
	ans := 0

	diff := make([]int, len(s))
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}

	left, sum := 0, 0
	for right, d := range diff {
		sum += d
		if sum > maxCost {
			sum -= diff[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	1004. 最大连续1的个数 III	https://leetcode.cn/problems/max-consecutive-ones-iii/
//输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//输出：6
//解释：[1,1,1,0,0,1,1,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 6。
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	ans := 0
	numZero := 0            //记录窗口里 0的数量
	for right < len(nums) { // 移动窗口右边
		if nums[right] == 0 {
			numZero++
		}
		for numZero > k { // 移动窗口左边
			if nums[left] == 0 {
				numZero--
			}
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
