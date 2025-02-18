package leetcode

func TopKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))

	for i := range len(nums) {
		if _, ok := counter[nums[i]]; !ok {
			counter[nums[i]] = 1
		} else {
			counter[nums[i]]++
		}
	}

	freq := make([][]int, len(nums)+1)

	for n, c := range counter {
		freq[c] = append(freq[c], n)
	}

	freqLength := len(freq)
	result := make([]int, 0, k)

outerLoop:
	for i := freqLength - 1; i >= 0; i-- {
		for j := range len(freq[i]) {
			result = append(result, freq[i][j])
			if len(result) == k {
				break outerLoop
			}
		}
	}

	return result
}
