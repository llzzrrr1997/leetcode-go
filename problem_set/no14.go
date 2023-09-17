package problemset

import "strings"

func longestCommonPrefix(strs []string) string {
	minStr := strs[0]
	for i := range strs {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}
	if minStr == "" {
		return ""
	}
	for i := 0; i < len(strs); {
		hasPre := strings.HasPrefix(strs[i], minStr)
		if !hasPre {
			i = 0
			minStr = minStr[:len(minStr)-1]
			continue
		}
		if hasPre && i == len(strs)-1 {
			return minStr
		}
		i++
	}
	return minStr
}
