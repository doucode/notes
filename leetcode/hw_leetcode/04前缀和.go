package hw_leetcode

// 4. 前缀和 总结：https://mp.weixin.qq.com/s/EwAH3JDs5WFO6-LFmI3-2Q
//代表题目：560🉑  触类旁通：523🉑.974🉑

//	560. 和为 K 的子数组	https://leetcode.cn/problems/subarray-sum-equals-k/
//	给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	mp := map[int]int{}
	mp[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre] += 1
	}
	return count
}

//	523. 连续的子数组和	https://leetcode.cn/problems/continuous-subarray-sum/
//给你一个整数数组 nums 和一个整数k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
//
//子数组大小 至少为 2 ，且
//子数组元素总和为 k 的倍数。
//如果存在，返回 true ；否则，返回 false 。
func checkSubarraySum(nums []int, k int) bool {
	m := map[int]int{0: -1} //-1代表从数组最开始
	sum := 0
	for i := range nums {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}
		if v, ok := m[sum]; ok {
			if i-v > 1 { //保证子数组最少2个元素
				return true
			}
		} else {
			m[sum] = i //余数为sum的下标
		}
	}
	return false
}

//	974. 和可被 K 整除的子数组	https://leetcode.cn/problems/subarray-sums-divisible-by-k/
//给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
//子数组 是数组的 连续 部分。
//输入：nums = [4,5,0,-2,-3,1], k = 5
//输出：7
func subarraysDivByK(nums []int, k int) int {
	record := map[int]int{0: 1} //初始化，记录record[0] = 1  考虑了前缀和本身被 k 整除的情况
	sum, ans := 0, 0
	for _, elem := range nums {
		sum += elem
		modulus := (sum%k + k) % k // go当被除数为负数时取模结果为负数，需要纠正
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}
