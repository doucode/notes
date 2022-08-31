package 算法思想

import (
	"container/heap"
	"math/rand"
	"sort"
)

//排序 1/4

//215. 数组中的第K个最大元素
//Input: [3,2,1,5,6,4] and k = 2
//Output: 5

//排序 ：时间复杂度 O(NlogN)，空间复杂度 O(1)
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//快速选择
func findKthLargest2(nums []int, k int) int {
	m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
	return selectSmallest(nums, 0, len(nums)-1, m)
}
func selectSmallest(nums []int, l, r, i int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums[q]
	}
	if i < k {
		return selectSmallest(nums, l, q-1, i)
	} else {
		return selectSmallest(nums, q+1, r, i-k)
	}
}
func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1) // 此处为优化，使得时间复杂度期望降为 O(n)，最坏时间复杂度为 O(n^2)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	// nums[l..i] <= nums[r]
	// nums[i+1..j-1] > nums[r]
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

//347. 前 K 个高频元素
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{key: key, count: count})
	}
	var result []int
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}

// Item define
type Item struct {
	key   int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	// 注意：因为golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

//451. 根据字符出现频率排序
//Input:
//"Aabb"
//
//Output:
//"bbAa"
//
//Explanation:
//"bbaA" is also a valid answer, but "Aabb" is incorrect.
//Note that 'A' and 'a' are treated as two different characters.
func frequencySort(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}

//75. 颜色分类
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
func frequencySort3(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
