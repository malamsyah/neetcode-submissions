func evalRPN(tokens []string) int {
	st := make([]int, 0)
	for _, t := range tokens {
		if t != "+" && t != "-" && t != "*" && t != "/" {
			val, _ := strconv.Atoi(t) 
			st = append(st, val)
			continue
		}

		a := st[len(st) - 2]
		b := st[len(st) - 1]
		if t == "+" {
			st[len(st) - 2] = a + b
		}
		if t == "-" {
			st[len(st) - 2] = a - b
		}
		if t == "*" {
			st[len(st) - 2] = a * b
		}
		if t == "/" {
			st[len(st) - 2] = a / b
		}

		st = st[:len(st) - 1]
	}

	return st[len(st) - 1]
}
