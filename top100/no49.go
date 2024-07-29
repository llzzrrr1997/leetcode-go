package top100

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, val := range strs {
		tmp := [26]byte{}
		for _, c := range val {
			tmp[c-'a']++
		}
		m[tmp] = append(m[tmp], val)
	}
	list := make([][]string, 0)
	for _, v := range m {
		list = append(list, v)
	}
	return list
}
