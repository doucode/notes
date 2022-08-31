package 算法思想

import "sort"

//贪心思想 1/11

//455. 分发饼干
//对每个孩子 i，都有一个胃口值g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j]。如果 s[j]>= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
//Input: grid[1,3], size[1,2,4]
//Output: 2
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}

//435. 无重叠区间
//题目描述：计算让一组区间不重叠所需要移除的区间个数。
//Input: [ [1,2], [1,2], [1,2] ]
//Output: 2

//Input: [ [1,2], [2,3] ]
//Output: 0
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}
