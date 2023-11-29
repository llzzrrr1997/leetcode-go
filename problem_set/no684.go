package problemset

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUF(n)
	var ret []int
	for _, edge := range edges {
		p := edge[0] - 1
		q := edge[1] - 1
		if uf.Connected(p, q) {
			ret = edge
		} else {
			uf.Union(q, p)
		}
	}
	return ret
}
