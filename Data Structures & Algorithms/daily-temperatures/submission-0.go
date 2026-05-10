func dailyTemperatures(temperatures []int) []int {
	arr := make([]int, 0, len(temperatures))
 	for i := 0; i < len(temperatures); i++ {
		curr := temperatures[i]
		ans := 0
	
		for j := i + 1; j < len(temperatures); j++ {
			num := temperatures[j]
			if curr < num {
				ans = j - i
				break
			}
		}

		arr = append(arr, ans)
 	}

	return arr
}
