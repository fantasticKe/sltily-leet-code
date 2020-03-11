package code_20200311

/**
难度: 简单
题目: 将数组分成和相等的三个部分

给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。

示例1：
输出：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

示例2：
输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false

示例3：
输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

提示:
3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4

*/

func CanThreePartsEqualSum(A []int) bool {
	sum := 0
	//求和
	for _, num := range A {
		sum += num
	}
	//判断数组元素和是否能被3整除
	if sum%3 != 0 {
		return false
	}
	//计算没部分数据的和
	sum /= 3
	temp := 0
	count := 0
	//遍历数组求和
	for i := 0; i < len(A)-1; i++ {
		temp += A[i]
		if temp == sum {
			count++
			if count == 2 {
				return true
			}
			temp = 0
		}
	}
	return false
}
