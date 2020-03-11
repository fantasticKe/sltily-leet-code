package code_20200311

import (
	"fmt"
	"testing"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	array1 := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	sum := CanThreePartsEqualSum(array1)
	array2 := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	notEqual := CanThreePartsEqualSum(array2)
	array3 := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	partsEqualSum := CanThreePartsEqualSum(array3)
	if !sum || notEqual || !partsEqualSum {
		t.Error("解答错误")
	}else {
		fmt.Println("解答正确")
	}
}
