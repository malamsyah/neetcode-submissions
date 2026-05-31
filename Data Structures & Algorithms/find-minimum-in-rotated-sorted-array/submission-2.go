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

func findMin(nums []int) int {
 l, r := 0, len(nums) - 1
 
 if nums[l] <= nums[r] {
    return nums[l]
 }

l = findFirst(l, r, func(x int) bool {
	return nums[x] <= nums[r]
})

 return nums[l]
}
