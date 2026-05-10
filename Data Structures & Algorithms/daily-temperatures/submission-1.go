type Stack struct {
	val int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
	arr := make([]int, len(temperatures))
	st := make([]Stack, 0)
	st = append(st, Stack{
			val: temperatures[0],
			idx: 0,
		})
 	for i := 1; i < len(temperatures); i++ {
		fmt.Println(st)
		curr := temperatures[i]
		for len(st) > 0 {
			top := st[len(st) - 1]
			if curr > top.val {
				arr[top.idx] = i - top.idx
				st = st[:len(st) - 1]
			} else {
				break
			}
		}

		st = append(st, Stack{
			val: curr,
			idx: i,
		})
 	}

	return arr
}
