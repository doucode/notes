package hw_leetcode

//并查集  总结：https://mp.weixin.qq.com/s/gUwLfi25TYamq8AJVIopfA
//代表题目：547🉑       触类旁通：684🉑.200*🆗.803

//	200. 岛屿数量	https://leetcode.cn/problems/number-of-islands/
//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]]
//输出：1
func numIslands(grid [][]byte) int {

}

//	547. 省份数量	https://leetcode.cn/problems/number-of-provinces/
//	有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
//	输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
//	输出：2
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int { // 查
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to) // 并
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}

	ans := 0
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return ans
}

//	684. 冗余连接	https://leetcode.cn/problems/redundant-connection/
//树可以看成是一个连通且 无环的无向图。
//
//给定往一棵n 个节点 (节点值1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges，edges[i] = [ai, bi]表示图中在 ai 和 bi 之间存在一条边。
//
//请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组edges中最后出现的边。
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1) // 树的结点的数量 = 边的数量 + 1
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

//803. 打砖块	https://leetcode.cn/problems/bricks-falling-when-hit/
func hitBricks(grid [][]int, hits [][]int) []int {

}
