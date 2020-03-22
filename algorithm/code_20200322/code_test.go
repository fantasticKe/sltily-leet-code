package code_20200322

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	if Reverse(123) == 321 && Reverse(-123) == -321 && Reverse(120) == 21 {
		println("解答正确")
	} else {
		println("解答错误")
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	if strings.EqualFold("fl", LongestCommonPrefix(strs1)) && strings.EqualFold("", LongestCommonPrefix(strs2)) {
		println("解答正确")
	} else {
		println("解答错误")
	}
}
