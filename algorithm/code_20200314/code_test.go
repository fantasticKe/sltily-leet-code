package code_20200314

import "testing"

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome(121) && !IsPalindrome(-121) && !IsPalindrome(10) {
		println("解答成功")
	} else {
		println("解答错误")
	}
}
