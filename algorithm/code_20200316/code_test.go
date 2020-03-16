package code_20200316

import "testing"

func TestCompressString(t *testing.T) {
	if CompressString("aabcccccaaa") == "a2b1c5a3" && CompressString("abbccd") == "abbccd" {
		println("解答正确")
	} else {
		println("解答错误")
	}
}
