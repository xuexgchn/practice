/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (55.37%)
 * Likes:    502
 * Dislikes: 0
 * Total Accepted:    80.6K
 * Total Submissions: 145.6K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 示例：
 * 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, top *ListNode

	// select list head
	if l1.Val < l2.Val {
		head = l1
		top = l1
		l1 = l1.Next
	} else {
		head = l2
		top = l2
		l2 = l2.Next
	}

	// loop
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			top.Next = l1
			l1 = l1.Next
		} else {
			top.Next = l2
			l2 = l2.Next
		}

		top = top.Next
	}

	if l1 != nil {
		top.Next = l1
	}
	if l2 != nil {
		top.Next = l2
	}

	return head
}

