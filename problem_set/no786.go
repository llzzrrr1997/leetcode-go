package problemset

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	type pair struct{ x, y int }
	list := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			list = append(list, pair{x: arr[i], y: arr[j]})
		}
	}
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		return a.x*b.y < a.y*b.x
	})
	return []int{list[k-1].x, list[k-1].y}
}
