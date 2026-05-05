func twoSum(nums []int, target int) []int {
    mn := make(map[int]int)

	for i, n := range nums {
		mn[n] = i
	}
	
	for i, n := range nums {
		j, found := mn[target - n]
		if found && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}
