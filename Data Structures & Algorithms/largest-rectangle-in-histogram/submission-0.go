func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	n := len(heights)
	mono := make([]int, 0, n)
	max := 0

	for i := 0; i < n; i++ {
		for len(mono) > 0 && heights[mono[len(mono)-1]] > heights[i] {
            top := mono[len(mono)-1]
            mono = mono[:len(mono)-1]

            var width int
            if len(mono) == 0 {
                width = i
            } else {
                width = i - mono[len(mono)-1] - 1
            }
            if area := heights[top] * width; area > max {
                max = area
            }
        }

        mono = append(mono, i)
	}

	return max
}
