package problemset

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	target := make([]byte, 0, len(words))
	for _, v := range words {
		if len(v) == 0 {
			return false
		}
		target = append(target, []byte(v)[0])
	}
	return string(target) == s
}
