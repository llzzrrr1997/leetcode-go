package problemset

import (
	"fmt"
	"testing"
)

func Test_minCommitter(t *testing.T) {
	ret := minCommitter([][]int{{0}, {5, 2, 3}, {0, 5}})
	fmt.Println(ret)
	ret = minCommitter([][]int{{1, 5, 2, 4}, {2, 9, 1}, {4}, {9}, {3, 1}})
	fmt.Println(ret)
}
