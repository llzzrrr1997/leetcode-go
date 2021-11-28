package problemset

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	targetMap := [26]int32{}
	for _, val := range p {
		targetMap[val-'a']++
	}
	curMap := [26]int32{}
	for i := 0; i < len(p); i++ {
		curMap[s[i]-'a']++
	}
	ret := []int{}
	if curMap == targetMap {
		ret = append(ret, 0)
	}
	n := len(p)
	for i := n; i < len(s); i++ {
		curMap[s[i]-'a']++
		curMap[s[i-n]-'a']--
		if curMap == targetMap {
			ret = append(ret, i-n+1)
		}
	}
	return ret
}
