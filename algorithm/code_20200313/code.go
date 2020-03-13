package code_20200313

import "sort"

/**
难度: 简单
题目：多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [3,2,3]
输出: 3

示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2

*/

//map法，遍历数据将数据作为key，出现次数作为value。
func MajorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	m := make(map[int]int, length)
	l := len(nums) / 2
	for _, num := range nums {
		m[num] = m[num] + 1
		if m[num] > l {
			return num
		}
	}
	return -1
}

//排序法，已知多数元素出现的次数大于len(nums)/2,则排序后的中间元素必定是多素元素
func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/**
难度: 中等
题目: 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归法
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	nextNode := AddTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{sum, nextNode}
	} else {
		tempNode := &ListNode{1, nil}
		return &ListNode{sum - 10, AddTwoNumbers(nextNode, tempNode)}
	}
}

//循环相加法
func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	next := &result
	num := 0
	for l1 != nil || l2 != nil || num > 0 {
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		*next = &ListNode{num % 10, nil}
		num /= 10
		next = &((*next).Next)
	}
	return result
}
