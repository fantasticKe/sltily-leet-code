package code_20200313

import (
	"strconv"
	"strings"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	arr1 := []int{3, 2, 3}
	arr2 := []int{2, 2, 1, 1, 1, 2, 2}
	element := MajorityElement(arr1)
	majorityElement := MajorityElement(arr2)
	if element == 3 && majorityElement == 2 {
		println("解答正确")
	} else {
		t.Fatal("解答错误")
	}
}

func TestMajorityElement2(t *testing.T) {
	arr1 := []int{3, 2, 3}
	arr2 := []int{2, 2, 1, 1, 1, 2, 2}
	element := MajorityElement2(arr1)
	majorityElement := MajorityElement2(arr2)
	if element == 3 && majorityElement == 2 {
		println("解答正确")
	} else {
		t.Fatal("解答错误")
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	numbers := AddTwoNumbers(l1, l2)
	i := strconv.Itoa(numbers.Val)
	for numbers.Next != nil {
		i = i + "->" + strconv.Itoa(numbers.Next.Val)
		numbers = numbers.Next
	}
	numbers2 := AddTwoNumbers2(l1, l2)
	ito := strconv.Itoa(numbers2.Val)
	for numbers2.Next != nil {
		ito = ito + "->" + strconv.Itoa(numbers2.Next.Val)
		numbers2 = numbers2.Next
	}
	println(i)
	println(ito)
	if strings.EqualFold("7->0->8", i) && strings.EqualFold("7->0->8", ito) {
		println("解答正确")
	} else {
		println("解答错误")
	}
}
