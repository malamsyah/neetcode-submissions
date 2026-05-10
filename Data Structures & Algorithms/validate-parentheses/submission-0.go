
func isValid(s string) bool {
	st := make([]rune, 0)
	for _, r := range s {
		if r == '{' || r == '(' || r == '[' {
			st = append(st, r)
		} else {
			if len(st) == 0 {
				return false
			}
			top := st[len(st) - 1]
			if r == '}' && top == '{' || r == ']' && top == '[' ||r == ')' && top == '(' {
				st = st[:len(st) - 1]
			} else {
				return false
			}
		}
	}

	return len(st) == 0
}


