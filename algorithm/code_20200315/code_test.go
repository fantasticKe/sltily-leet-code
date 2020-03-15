package code_20200315

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	if LengthOfLongestSubstring(s1) == 3 && LengthOfLongestSubstring(s2) == 1 && LengthOfLongestSubstring(s3) == 3 {
		println("解法1解答正确")
	} else {
		println("解法1解答错误")
	}

	if LengthOfLongestSubstring2(s1) == 3 && LengthOfLongestSubstring2(s2) == 1 && LengthOfLongestSubstring2(s3) == 3 {
		println("解法2解答正确")
	} else {
		println("解法2解答错误")
	}
}
