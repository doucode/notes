package hw_leetcode

import "sort"

//12. 回溯 总结：https://mp.weixin.qq.com/s/nMUHqvwzG2LmWA9jMIHwQQ
//代表题目：698🉑 触类旁通：577🔒

//	698. 划分为k个相等的子集	https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
//给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
func canPartitionKSubsets(nums []int, k int) bool {
	n = len(nums)
	// 桶比数多，无解
	if k > n {
		return false
	}
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 不能整除，无解
	if sum%k > 0 {
		return false
	}
	target = sum / k
	// 倒序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	// 单个数比目标值大，无解
	if nums[0] > target {
		return false
	}
	used := make([]bool, n)
	return traverse(nums, 0, k, 0, used)
}

var n int
var target int

// 回溯算法 Bucket表示当前桶内的和；k表示第几个桶；start表示当前数的索引；used表示访问记录
func traverse(nums []int, bucket, k, start int, used []bool) (ok bool) {
	// base case.如果 已经到达最后一个桶 则必然满足条件（总和能等分，且前面都满足条件）
	if k == 0 {
		return true
	}
	// 回溯：如果当前桶的和已经等于target 则进入下一个桶
	if bucket == target {
		// 装满了当前桶，递归穷举下一个桶的选择
		// 让下一个桶从 nums[0] 开始选数字
		return traverse(nums, 0, k-1, 0, used)
	}
	// 遍历数组其他数字
	for i := start; i < n; i++ {
		// 访问记录，剪枝
		if used[i] {
			continue
		}
		// 桶装不下当前这个数字，剪枝
		if bucket+nums[i] > target {
			continue
		}
		//前序计算 桶装进这个数子后求和
		bucket += nums[i]
		used[i] = true
		// 进入下一层 数字回溯。如果下一层依旧满足则继续回溯，直到base case
		if traverse(nums, bucket, k, i+1, used) {
			return true
		}
		// 如果 后续层（数字）不满足条件则撤销选择。进入兄弟节点重新回溯。
		bucket -= nums[i]
		used[i] = false
	}
	// base case 如果所有数字都不满足该桶条件 则直接返回无解
	return false
}
