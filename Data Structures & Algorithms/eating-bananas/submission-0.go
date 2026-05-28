func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	
	ans := r

	for l <= r {
		mid := (l + r) / 2
		currH := 0
		for _, p := range piles {
			currH += int(math.Ceil(float64(p) / float64(mid)))
		}

		if currH <= h {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return ans
}

// 1, 2, 3, 4, 5
