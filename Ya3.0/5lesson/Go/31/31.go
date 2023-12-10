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
	for i := 1; i != n + 1; i++ {
		g[i] = make([]int, 0)
	}

	for i := 0; i != m; i++ {
		var v1, v2 int
		fmt.Fscan(cin, &v1, &v2)

		g[v1] = append(g[v1], v2)
		g[v2] = append(g[v2], v1)
	}

	visited := make([]bool, n + 1)
	cnt := 0
	dfs(visited, g, 1, &cnt)

	fmt.Fprintln(cout, cnt)
	for i, st := range visited {
		if st {
			fmt.Fprintf(cout, "%d ", i)
		}
	}

	fmt.Fprintln(cout)
	cout.Flush()
}

func dfs(visited []bool, g [][]int, x int, cnt *int) {
	visited[x] = true
	*cnt++
	for _, v := range g[x] {
		if !visited[v] {
			dfs(visited, g, v, cnt)
		}
	}
}