package getting_started

//no387
func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

//no383
func canConstruct(ransomNote string, magazine string) bool {
	rMap := make([]int, 26)
	mMap := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		rMap[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if rMap[i] > mMap[i] {
			return false
		}
	}
	return true
}

//no242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		m[t[i]-'a']--
		if m[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
