package leetcode

func ContainsDuplicate(nums []int) bool {
	tracker := make(map[int]bool, 10)

	for i := 0; i < len(nums); i++ {
		if _, ok := tracker[nums[i]]; ok {
			return true
		} else {
			tracker[nums[i]] = true
		}
	}

	return false
}
