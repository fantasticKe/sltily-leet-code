package code_20200322

import (
	"math"
	"strings"
)

/**
难度：简单
题目：整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func Reverse(x int) int {
	min := math.MinInt32 / 10
	max := math.MaxInt32 / 10
	y := 0
	for x != 0 {
		if y < min || y > max {
			return 0
		}
		y = y*10 + x%10
		x = x / 10
	}
	return y
}

/**
难度： 简单
题目： 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:
所有输入只包含小写字母 a-z 。
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minIndex := 0
	for i, s := range strs {
		if len(strs[minIndex]) > len(s) {
			minIndex = i
		}
	}
	for i := len(strs[minIndex]); i > 0; i-- {
		for x, s := range strs {
			if strings.HasPrefix(s, strs[minIndex][:i]) {
				if x == len(strs)-1 {
					return strs[minIndex][:i]
				}
				continue
			} else {
				break
			}
		}
	}
	return ""
}
