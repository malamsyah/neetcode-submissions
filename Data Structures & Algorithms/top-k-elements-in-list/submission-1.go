func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v] += 1
	}

	buc := make([][]int, len(nums) + 1)
	for k, v := range freq {
		buc[v] = append(buc[v], k)
	}

	ans := make([]int, 0)
	for i := len(buc) - 1; len(ans) < k; i-- {
		for _ , v := range buc[i] {
			if len(ans) == k {
				break
			}
			ans = append(ans, v)
		}
	}

	return ans
}
