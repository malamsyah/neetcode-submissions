func threeSum(nums []int) [][]int {
	ans := make(map[[3]int]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				if i == j || i == k || j == k {
					continue
				}

				ni := nums[i]
				nj := nums[j]
				nk := nums[k]

				if ni + nj + nk == 0{
					ans[sortSet(ni, nj, nk)] = struct{}{}
				}
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
