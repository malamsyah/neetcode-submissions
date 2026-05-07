func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// fmt.Println(nums)
	ans := make(map[[3]int]struct{})
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l == i {
				l++
				continue
			}

			if r == i {
				r--
				continue
			}

			curr := target + nums[l] + nums[r]
			if curr == 0 {
				ans[sortSet(target, nums[l], nums[r])] = struct{}{}
				l++
				r--
			} else if curr > 0 {
				r--
			} else if curr < 0 {
				l++
			}
		}
	}

	tmp := make([][]int, len(ans))
	i := 0
	for k, _ := range ans {
		tmp[i] = []int{k[0], k[1], k[2]}
		i++
	}

	return tmp
}

func sortSet(i, j, k int)[3]int{
	if i > j {
		i, j = j, i
	}
	if i > k {
		i, k = k, i
	}
	if j > k {
		j, k = k, j
	}

	return [3]int{i, j, k}
}
