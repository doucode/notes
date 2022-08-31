package hw_leetcode

//8. DFS 总结：https://mp.weixin.qq.com/s/IZQkb-M27dt-AZ1VICThOw
//代表题目：200🉑 触类旁通：1254🉑

//	200. 岛屿数量	https://leetcode.cn/problems/number-of-islands/
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]]
//输出：1
func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	h, w := len(grid), len(grid[0])
	if r < 0 || r >= h || c < 0 || c >= w {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

//	1254. 统计封闭岛屿的数目	https://leetcode.cn/problems/number-of-closed-islands/
//二维矩阵 grid由 0（土地）和 1（水）组成。岛是由最大的4个方向连通的 0组成的群，封闭岛是一个完全 由1包围（左、上、右、下）的岛。
//
//请返回 封闭岛屿 的数目。
//输入：grid = [[1,1,1,1,1,1,1,0],
//			  [1,0,0,0,0,1,1,0],
//			  [1,0,1,0,1,1,1,0],
//			  [1,0,0,0,0,1,0,1],
//			  [1,1,1,1,1,1,1,0]]
//输出：2
func closedIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 1 {
			return
		}
		//递归出口  已经遍历过(i,j)
		if visited[i][j] {
			return
		}
		//遍历位置(i,j)
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	//首先把四周给遍历了
	for j := 0; j < n; j++ {
		dfs(0, j)   //最上面
		dfs(m-1, j) //最下面
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)   //最左边
		dfs(i, n-1) //最右边
	}
	//遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == 0 && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
