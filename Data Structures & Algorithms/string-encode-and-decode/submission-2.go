type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "0#" {
		return []string{""}
	}

	ans := make([]string, 0)
	n := len(encoded)
	i := 0
	for i < n {
		j := i
		for encoded[j] != '#' {
			j++
		}

		subN, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			return nil // malformed input
		}

		start := j + 1
		end := start + subN
		ans = append(ans, encoded[start:end])
		i = end
	}

	return ans
}
