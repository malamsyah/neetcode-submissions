func isAnagram(s string, t string) bool {
	ns, nt := len(s), len(t)

	if ns != nt {
		return false
	}

	ms := make(map[rune]int, ns)
	mt := make(map[rune]int, nt)

	for _, c := range s {
		ms[c] += 1
	}

	for _, c := range t {
		mt[c] += 1
	}

	if len(ms) != len(mt) {
		return false
	}

	for k, v := range ms {
		x, found := mt[k]
		if !found {
			return false
		}

		if x != v {
			return false
		}
	}

	return true
}
