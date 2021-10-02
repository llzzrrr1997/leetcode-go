package foundation

import (
	"fmt"
	"strings"
)

//no49
func groupAnagrams(strs []string) [][]string {
	retM := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		tmp := [26]int{}
		for _, val := range strs[i] {
			tmp[val-'a']++
		}
		retM[tmp] = append(retM[tmp], strs[i])
	}
	ret := make([][]string, 0, len(retM))
	for _, val := range retM {
		ret = append(ret, val)
	}
	return ret
}

//no43
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num2) < len(num1) {
		num1, num2 = num2, num1
	}
	tepList := make([]string, 0, len(num1))
	zeroCnt := 0
	for i := len(num1) - 1; i >= 0; i-- {
		zeroS := strings.Repeat("0", zeroCnt)
		zeroCnt++
		if num1[i] == '0' {
			tepList = append(tepList, "0")
		}
		//乘法
		more := byte(0)
		tmpS := make([]byte, 0)
		for j := len(num2) - 1; j >= 0; j-- {
			tmp := more + (num1[i]-'0')*(num2[j]-'0')
			tmpS = append(tmpS, tmp%10+'0')
			more = tmp / 10
		}
		if more != 0 {
			tmpS = append(tmpS, more+'0')
		}
		left, right := 0, len(tmpS)-1
		for left < right {
			tmpS[left], tmpS[right] = tmpS[right], tmpS[left]
			left++
			right--
		}
		//补0
		tepList = append(tepList, string(tmpS)+zeroS)
	}
	fmt.Println(tepList)
	ret := "0"
	for i := 0; i < len(tepList); i++ {
		ret = addStrings2(ret, tepList[i])
	}
	return ret
}

func addStrings2(num1 string, num2 string) string {
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
