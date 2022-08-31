package hw_leetcode

import "sort"

//10. 贪心算法 总结：https://mp.weixin.qq.com/s/k-z_oewAqMYc3vpmOm4gEQ
//代表题目：452 触类旁通：1247.1231🔒.45🉑 .621.376

//	45. 跳跃游戏 II	https://leetcode.cn/problems/jump-game-ii/
//给你一个非负整数数组nums ，你最初位于数组的第一个位置。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//假设你总是可以到达数组的最后一个位置

//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
func jump(nums []int) int {
	step, end, maxPosition := 0, 0, 0 // end是当前能跳的右边界
	i := 0
	for end < len(nums)-1 {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
		i++
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	452. 用最少数量的箭引爆气球	https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
//输入：points = [[10,16],[2,8],[1,6],[7,12]]
//输出：2
//解释：气球可以用2支箭来爆破:
//-在x = 6处射出箭，击破气球[2,8]和[1,6]。
//-在x = 11处发射箭，击破气球[10,16]和[7,12]。、
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}
	// 对原数组进行排序，得到升序排列数组
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	arrows := 1
	for i := 1; i < n; i++ {
		// 两个气球不重叠，则需要的弓箭数+1
		if points[i][0] > points[i-1][1] {
			arrows++
		} else {
			// 两个气球重叠，则更新右边气球的最小右边界
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return arrows
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//	376. 摆动序列	https://leetcode.cn/problems/wiggle-subsequence/
//给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
//输入：nums = [1,7,4,9,2,5]
//输出：6
//解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

//	621. 任务调度器	https://leetcode.cn/problems/task-scheduler/
//给你一个用字符数组tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
//
//然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//
//你需要计算完成所有任务所需要的 最短时间 。
//输入：tasks = ["A","A","A","B","B","B"], n = 2
//输出：8
//解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	list := make([]int, 26)
	for k := range tasks {
		list[tasks[k]-'A']++
	}
	maxExec, maxCount := 0, 0
	for k := range list {
		if list[k] > maxExec {
			maxExec, maxCount = list[k], 1
		} else if list[k] == maxExec {
			maxCount++
		}
	}
	time := (maxExec-1)*(n+1) + maxCount
	if time < len(tasks) {
		return len(tasks)
	}
	return time
}

//	1247. 交换字符使得字符串相同	https://leetcode.cn/problems/minimum-swaps-to-make-strings-equal/
//有两个长度相同的字符串 s1 和 s2，且它们其中 只含有 字符 "x" 和 "y"，你需要通过「交换字符」的方式使这两个字符串相同。
//输入：s1 = "xx", s2 = "yy"
//输出：1
//解释：
//交换 s1[0] 和 s2[1]，得到 s1 = "yx"，s2 = "yx"。
func minimumSwap(s1 string, s2 string) int {
	x, y := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			x++
		} else {
			y++
		}
	}
	if x%2+y%2 == 1 {
		return -1
	}
	return x/2 + y/2 + 2*(x%2)
}
