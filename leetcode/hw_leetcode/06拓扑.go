package hw_leetcode

//6. 拓扑 总结：https://mp.weixin.qq.com/s/7nP92FhCTpTKIAplj_xWpA
//代表题目：210🉑 触类旁通：269🔒

//	210. 课程表 II	https://leetcode.cn/problems/course-schedule-ii/
//现在你总共有 numCourses 门课需要选，记为0到numCourses - 1。给你一个数组prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修bi 。
//
//例如，想要学习课程 0 ，你需要先完成课程1 ，我们用一个匹配来表示：[0,1] 。
//返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

//输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//输出：[0,2,1,3]
//解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
//因此，一个正确的课程顺序是[0,1,2,3] 。另一个正确的排序[0,2,1,3] 。

func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses) // 入度
	frees := make([][]int, numCourses)
	next := make([]int, 0, numCourses)
	for _, v := range prerequisites {
		in[v[0]]++                              // 累加 入度
		frees[v[1]] = append(frees[v[1]], v[0]) // 指向的结点数组
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			next = append(next, i) //入度为0的结点先写进答案
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--         // 入度 -1
			if in[vv] == 0 { // 入度为0 写进答案
				next = append(next, vv)
			}
		}
	}
	if len(next) == numCourses { // next 和 numCourses长度相等则输出
		return next
	}
	return []int{}
}
