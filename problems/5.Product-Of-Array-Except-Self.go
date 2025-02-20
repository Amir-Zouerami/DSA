package leetcode

func ProductExceptSelf(nums []int) []int {
	numsLength := len(nums)
	result := make([]int, numsLength)

	prefix := 1
	for i := 0; i < numsLength; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := numsLength - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
