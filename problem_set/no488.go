package problemset

import "fmt"

func findMinStep(board string, hand string) int {
	q := make([][]string, 0)
	q = append(q, []string{board, hand})
	ret := 0
	m := make(map[string]bool)
	for len(q) > 0 {
		ret++
		n := len(q)
		for n > 0 {
			n--
			cur := q[0]
			q = q[1:]
			b, h := cur[0], cur[1]
			//暴力模拟
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(h); j++ {
					b2 := del(b[0:i] + string(h[j]) + b[i:])
					h2 := h[0:j] + h[j+1:]
					if b2 == "" {
						return ret
					}
					key := fmt.Sprintf("%s_%s", b2, h2)
					if m[key] {
						continue
					}
					m[key] = true
					q = append(q, []string{b2, h2})
				}
			}
		}
	}
	return -1
}

//递归碰撞删除所有长度3及以上的子串
func del(s string) string {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				return del(s[0:i-count] + s[i:])
			}
			count = 1
		}
	}
	if count >= 3 {
		return del(s[0 : len(s)-count])
	}
	return s
}
