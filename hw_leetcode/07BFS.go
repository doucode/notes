package hw_leetcode

import "math"

//7. BFS 总结：https://mp.weixin.qq.com/s/WH_XGm1-w5882PnenymZ7g
//代表题目：127🉑 触类旁通：139🉑.317🔒.815*🉑

//	139. 单词拆分	https://leetcode.cn/problems/word-break/
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	queue := []int{}
	queue = append(queue, 0) // 记录开始下标，从0开始
	visited := map[int]bool{}

	for len(queue) != 0 {
		start := queue[0]
		queue = queue[1:] //	出列
		if visited[start] {
			continue
		}
		visited[start] = true
		for i := start + 1; i <= n; i++ {
			prefix := s[start:i]
			if wordMap[prefix] {
				if i < n {
					queue = append(queue, i)
				} else {
					return true //	s被全部遍历
				}
			}
		}
	}
	return false
}

//	127. 单词接龙	https://leetcode.cn/problems/word-ladder/
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
//输出：5
//解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

// 	815. 公交路线	https://leetcode.cn/problems/bus-routes/
//给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
//
//例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
//现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
//
//求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
//输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
//输出：2
//解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
func numBusesToDestination(routes [][]int, source int, target int) int {
	// 出发站就是目的站
	if source == target {
		return 0
	}

	stations := map[int][]int{}
	for i, route := range routes {
		for _, station := range route {
			stations[station] = append(stations[station], i) //记录每个车站的所属路线
		}
	}

	type pair struct {
		id, step int
	}

	visited := make([]bool, len(routes)) //记录 线路 是否遍历过
	q := []pair{{source, 0}}
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			p := q[0]
			q = q[1:]
			for _, route := range stations[p.id] {
				if !visited[route] {
					visited[route] = true
					for _, station := range routes[route] {
						if station == target {
							return p.step + 1
						}
						q = append(q, pair{station, p.step + 1})
					}
				}
			}
		}
	}

	return -1
}

//
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	routesNum := map[int][]int{}
	for i, r := range routes {
		for _, s := range r {
			routesNum[s] = append(routesNum[s], i)
		}
	}

	visited := make([]bool, len(routes))

	type pair struct {
		id   int
		step int
	}

	q := []pair{{source, 0}}

	for len(q) != 0 {
		start := q[0]
		q = q[1:]
		for _, route := range routesNum[start.id] {
			if !visited[route] {
				visited[route] = true
				for _, st := range routes[route] {
					if st == target {
						return start.step + 1
					} else {
						q = append(q, pair{st, start.step + 1})
					}
				}
			}
		}
	}

}

//
