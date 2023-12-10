package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, m int
	fmt.Fscan(cin, &n, &m)

	g := make([][]int, n + 1)
	for i := 0; i != n + 1; i++ {
		g[i] = make([]int, 0)
	}

	for i := 0; i != m; i++ {
		var v1, v2 int
		fmt.Fscan(cin, &v1, &v2)

		g[v1] = append(g[v1], v2)
		g[v2] = append(g[v2], v1)
	}

	colors := make([]int, n + 1)
	ans := true
	for v := 1; v != n + 1; v++ {
		if colors[v] == 0 {
			colors[v] = 1
			ans = ans && dfs(g, colors, v)
		}
	}

	if ans {
		fmt.Fprintln(cout, "YES")
	} else {
		fmt.Fprintln(cout, "NO")
	}

	cout.Flush()
}

func dfs(g [][]int, colors []int, x int) bool {
	for _, v := range g[x] {
		switch colors[v] {
		case colors[x]:
			return false
		case 0:
			colors[v] = 3 - colors[x]
			if !dfs(g, colors, v) {
				return false
			}
		}
	}
	return true
}