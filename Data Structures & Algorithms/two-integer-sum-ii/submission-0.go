func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		curr := numbers[l] + numbers[r]
		if curr == target {
			return []int{l+1, r+1}
		}
		if curr < target {
			l++
		}
		if curr > target {
			r--
		}
	}

	return []int{}
}
