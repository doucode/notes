package 算法思想

import (
	"math"
	"sort"
	"strings"
)

//双指针 7/7

//167. 两数之和 II - 输入有序数组
//Input: numbers={2, 7, 11, 15}, target=9
//Output: index1=1, index2=2
//题目描述：在 有序 数组中找出两个数，使它们的和为 target。
func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1 //双指针
	for lo < hi {
		sum := numbers[lo] + numbers[hi] //相加
		if sum == target {
			return []int{lo + 1, hi + 1}
		} else if sum < target { //sum小于target，lo指针+1
			lo++
		} else { //sum大于target，hi指针-1
			hi--
		}
	}
	return []int{-1, -1}
}

//633. 平方数之和
//     Input: 5
//     Output: True
//     Explanation: 1 * 1 + 2 * 2 = 5
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		b := math.Sqrt(float64(c - a*a)) //先强转成float64 再开平方
		if b == math.Floor(b) {          //取b float64 的最大整数
			return true
		}
	}
	return false
}

//345. 反转字符串中的元音字母
//示例 ：
//输入：s = "hello"
//输出："holle"
func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && strings.Contains("aeiouAEIOU", string(t[i])) { //strings.Contains(string,substr)返回bool
			i++ //包含返回true
		}
		for j > 0 && strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}

//680. 验证回文字符串 Ⅱ
//给定一个非空字符串s，最多删除一个字符。判断是否能成为回文字符串。
//示例 1:
//输入: s = "aba"
//输出: true
//示例 2:
//输入: s = "abca"
//输出: true
//解释: 你可以删除c字符。
func validPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] == s[hi] {
			lo++
			hi--
		} else {
			flag1, flag2 := true, true
			for i, j := lo, hi-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
				}
			}
			for i, j := lo+1, hi; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

//88. 合并两个有序数组
//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]

//解法1
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

//解法2
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

//141. 环形链表
//判断链表是否存在环
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	bike, car := head, head.Next
	for bike != nil && car != nil && car.Next != nil {
		if car == bike {
			return true
		}
		car = car.Next.Next
		bike = bike.Next
	}
	return false
}

//524. 通过删除字母匹配到字典里最长单词
//示例 1：
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
func findLongestWord(s string, d []string) string {
	res := ""
	for i := 0; i < len(d); i++ {
		ps, pd := 0, 0
		for ps < len(s) && pd < len(d[i]) {
			if s[ps] == d[i][pd] {
				pd++
			}
			ps++
		}
		if pd == len(d[i]) && (pd > len(res) || pd == len(res) && res < d[i]) {
			res = d[i]
		}
	}
	return res
}
