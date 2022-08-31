package 数据结构相关

// 链表 10/10

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}
	return l1
}

// 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 21. 合并两个有序链表  https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2) // 递归
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1) // 递归
		return list2
	}
}

// 83. 删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

//19. 删除链表的倒数第 N 个结点  https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

// 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{0, head}
	temp := dummyNode
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyNode.Next
}

// 445. 两数相加 II	https://leetcode.cn/problems/add-two-numbers-ii/description/
//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 = []int{}, []int{} //  栈
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var res *ListNode
	var cur = 0
	for len(s1) > 0 || len(s2) > 0 || cur != 0 {
		if len(s1) > 0 {
			cur += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			cur += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		node := &ListNode{Val: cur % 10}
		cur /= 10
		node.Next = res
		res = node
	}
	return res
}

//	234. 回文链表	https://leetcode.cn/problems/palindrome-linked-list/description/
func isPalindrome(head *ListNode) bool {
	slice := []int{}
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}

// 725. 分隔链表
//题目描述：把链表分隔成 k 部分，每部分的长度都应该尽可能相同，
//排在前面的长度应该大于等于后面的。
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	quotient, remainder := n/k, n%k

	parts := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient
		if i < remainder {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}

// 328. 奇偶链表
//给定单链表的头节点head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
//
//第一个节点的索引被认为是 奇数 ， 第二个节点的索引为偶数 ，以此类推。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
