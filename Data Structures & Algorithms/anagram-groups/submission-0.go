func groupAnagrams(strs []string) [][]string {
	hashGroup := make(map[int][]string)

	for _, v := range strs {
		h := hashStr(v)
		if hashGroup[h] == nil {
			hashGroup[h] = []string{v}
		} else {
			hashGroup[h] = append(hashGroup[h], v)
		}
	}

	res := make([][]string, len(hashGroup))
	i := 0
	for _, v := range hashGroup {
		res[i] = v
		i++
	}

	return res
}

// hashStr -> "act" or "cat" should return the same hash
func hashStr(str string) int {
	n := len(str) * 7 % math.MaxInt
	sum := 1
	for _, v := range str {
		sum *= int(v) % math.MaxInt
	}

	return sum * n % math.MaxInt
}
