func findFirst(l, h int, pred func(x int) bool) int {
	for l < h {
		m := l + (h - l) / 2
		if pred(m) {
			h = m
		} else {
			l = m + 1
		}
	}

	return l // l == h
}

func search(nums []int, target int) int {
	l, h := 0, len(nums) - 1

	p := findFirst(l, h, func(x int) bool {
		return nums[x] <= nums[h]
	})

	l, h = p, len(nums) - 1
	
	if target > nums[len(nums) - 1] {
		l, h = 0, p - 1
	}

	if l > h {
		return -1
	}
	
	i := findFirst(l, h + 1, func(x int) bool {
		return nums[x] >= target
	})

	if i <= h && nums[i] == target {
		return i
	} else {
		return -1
	}
}
