package foundation

//no415
func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	more := byte(0)
	i, j := n1-1, n2-1
	ret := make([]byte, 0)
	for i >= 0 || j >= 0 {
		tmp := more
		if i >= 0 {
			tmp += num1[i] - '0'
			i--
		}
		if j >= 0 {
			tmp += num2[j] - '0'
			j--
		}
		ret = append(ret, tmp%10+'0')
		more = tmp / 10
	}
	if more != 0 {
		ret = append(ret, more+'0')
	}
	left, right := 0, len(ret)-1
	for left < right {
		ret[left], ret[right] = ret[right], ret[left]
		left++
		right--
	}
	return string(ret)
}

//409
func longestPalindrome(s string) int {
	cntMap := make(map[int32]int)
	for _, val := range s {
		cntMap[val]++
	}
	ret, hasOdd := 0, false
	for _, val := range cntMap {
		if val%2 == 1 {
			ret += val - 1
			hasOdd = true
		} else {
			ret += val
		}
	}
	if hasOdd {
		ret++
	}
	return ret
}
