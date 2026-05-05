func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	n := len(nums)

	// Left
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i]
	}

	// Right

	right[n - 1] = nums[n - 1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i + 1] * nums[i]
	}

	// fmt.Println(left)
	// fmt.Println(right)

	ans := make([]int, len(nums))
	ans[0] = right[1]
	ans[n-1] = left[n - 2]
	for i := 1; i < n - 1; i++ {
		ans[i] = left[i-1] * right[i + 1]
	}

	return ans
}
