package leetcode

func LongestConsecutive(nums []int) int {
	hash := make(map[int]struct{})

	for i := range nums {
		hash[nums[i]] = struct{}{}
	}

	longest := 0

	for num := range hash {
		if _, ok := hash[num-1]; !ok {
			length := 1

			for {
				if _, ok := hash[num+length]; ok {
					length++
				} else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

/*
  Note: You could use sorting to solve this but
  the time complexity would be O(n * logn) which is not great.
*/
