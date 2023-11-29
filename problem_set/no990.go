package problemset

type UF struct {
	// 记录连通分量
	count int
	// 节点 x 的父节点是 parent[x]
	parent []int
}

func NewUF(n int) *UF {
	uf := &UF{count: n, parent: make([]int, n)}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (u *UF) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

// 将节点 p 和节点 q 连通
func (u *UF) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	u.parent[rootQ] = rootP
	// 两个连通分量合并成一个连通分量
	u.count--
}

// 判断节点 p 和节点 q 是否连通
func (u *UF) Connected(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	return rootP == rootQ
}

// 返回图中的连通分量个数
func (u *UF) Count() int {
	return u.count
}

func equationsPossible(equations []string) bool {
	uf := NewUF(26)
	for _, equation := range equations {
		if equation[1] == '=' {
			p := equation[0] - 'a'
			q := equation[3] - 'a'
			uf.Union(int(p), int(q))
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			p := equation[0] - 'a'
			q := equation[3] - 'a'
			if uf.Connected(int(p), int(q)) {
				return false
			}
		}
	}
	return true
}
