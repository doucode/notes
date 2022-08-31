package hw_leetcode

//13. 栈与队列 总结：https://mp.weixin.qq.com/s/xBcHyvHlWq4P13fzxEtkPg
//代表题目：21🉑

type ListNode struct {
	Val  int
	Next *ListNode
}

//	21. 合并两个有序链表	https://leetcode.cn/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}
