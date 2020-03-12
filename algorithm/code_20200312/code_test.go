package code_20200312

import (
	"strings"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	str1, str2 := "ABCABC", "ABC"
	s := GcdOfStrings(str1, str2)
	str3, str4 := "ABABAB", "ABAB"
	s1 := GcdOfStrings(str3, str4)
	str5, str6 := "LEET", "CODE"
	s2 := GcdOfStrings(str5, str6)
	if !strings.EqualFold("ABC", s) || !strings.EqualFold("AB", s1) || !strings.EqualFold("", s2) {
		t.Fatal("解答错误")
	} else {
		println("解答正确")
	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := TwoSum(nums, target)
	target1 := 10
	twoSum := TwoSum(nums, target1)
	if sum[0] == 0 && sum[1] == 1 && len(twoSum) == 0 {
		println("解答正确")
	} else {
		t.Fatal("解答错误")
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := TwoSum2(nums, target)
	target1 := 10
	twoSum := TwoSum2(nums, target1)
	if sum[0] == 0 && sum[1] == 1 && len(twoSum) == 0 {
		println("解答正确")
	} else {
		t.Fatal("解答错误")
	}
}
