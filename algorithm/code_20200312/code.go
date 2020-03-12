package code_20200312

import "strings"

/**
难度： 简单
题目：字符串的最大公因子
对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。

示例 1：
输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"

示例 2：
输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"

示例 3：
输入：str1 = "LEET", str2 = "CODE"
输出：""


提示：

1 <= str1.length <= 1000
1 <= str2.length <= 1000
str1[i] 和 str2[i] 为大写英文字母

*/

func GcdOfStrings(str1 string, str2 string) string {
	if !strings.EqualFold(str1+str2, str2+str1) {
		return ""
	}
	i := gcd(len(str1), len(str2))
	return str1[0:i]
}

//辗转相除
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/**
难度：简单
题目：两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/

func TwoSum(nums []int, target int) []int {
	//目标元素依次相减得到的数据是否存在于数组当中
	for i, num := range nums {
		temp := target - num
		b, j := contains(nums, temp)
		if b && i != j {
			return []int{i, j}
		}
	}
	return nil
}

func contains(nums []int, target int) (bool, int) {
	for i, num := range nums {
		if num == target {
			return true, i
		}
	}
	return false, -1
}

func TwoSum2(nums []int, target int) []int {
	//map法
	m := make(map[int]int)
	for i, num := range nums {
		if k, ok := m[target-num]; ok {
			return []int{k, i}
		}
		m[num] = i
	}
	return nil
}
