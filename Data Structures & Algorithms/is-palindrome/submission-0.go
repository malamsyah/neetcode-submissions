func isPalindrome(s string) bool {
	i, j := 0, len(s) -1
	n := len(s)
	for i < j {
		var x, y byte
		var valid bool
		for i < n {
			if x, valid = isChar(s[i]); !valid {
				i++
			} else {
				break
			}
		}
		for j >= 0 {
			// fmt.Println("j: ", j)
			if y, valid = isChar(s[j]); !valid {
				j--
			} else {
				break
			}
		}
		// fmt.Println(rune(x), " ", rune(y))
		if x != y {
			return false
		}
		i++
		j--
	}

	return true
}

func isChar(s byte) (byte, bool) {
	if s >= '0' && s <= '9' {
		return s, true
	}
	if s >= 'a' && s <= 'z' {
		return s - 'a', true
	} 
	if s >= 'A' && s <= 'Z' {
		l := s - 'A'
		return l, true
	}

	return 0, false
}
