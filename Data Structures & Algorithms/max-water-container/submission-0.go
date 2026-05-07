func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	ans := -1
	for l < r {
		curr := min(heights[l], heights[r]) * (r - l)
		if curr > ans {
			ans = curr
		}

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return ans
}

func min(i, j int) int{
	if i > j {
		return j
	}
	return i
}
