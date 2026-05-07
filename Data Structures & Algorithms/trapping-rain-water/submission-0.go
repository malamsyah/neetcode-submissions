func trap(heights []int) int {
	n := len(heights)
	maxL := make([]int, n)
	maxR := make([]int, n)

	maxL[0] = heights[0]
	for i := 1; i < n; i++{
		if heights[i] > maxL[i - 1] {
			maxL[i] = heights[i] 
		} else {
			maxL[i] = maxL[i - 1]
		}
	}

	maxR[n - 1] = heights[n - 1]
	for i := n - 2; i >= 0; i--{
		if heights[i] > maxR[i + 1] {
			maxR[i] = heights[i] 
		} else {
			maxR[i] = maxR[i + 1]
		}
	}

	// fmt.Println(maxL)
	// fmt.Println(maxR)
	ans := 0
	for i := 1; i < n - 1; i++ {
		h := heights[i]
		if h < maxL[i-1] && h < maxR[i+1] {
			ans += (min(maxL[i-1], maxR[i+1]) - h)
		} 
	}

	return ans
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
