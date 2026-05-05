type elem struct {
	n int
	num int
}

func topKFrequent(nums []int, k int) []int {
	elemMap := map[int]*elem{}

	for _, num := range nums {
		if e, ok := elemMap[num]; ok {
            e.n++
        } else {
            elemMap[num] = &elem{n: 1, num: num}
        }
	}

	elemArr := make([]*elem, 0, len(elemMap))
    for _, v := range elemMap {
        elemArr = append(elemArr, v)
    }

	sort.Slice(elemArr, func(i, j int) bool {
        return elemArr[i].n > elemArr[j].n
    })

    ans := make([]int, 0, k)
    for r := 0; r < k; r++ {
        ans = append(ans, elemArr[r].num)
    }

    return ans
}
