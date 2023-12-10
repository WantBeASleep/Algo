package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	cc := 0
	ans := make([][]int, 0, n)
	for i := 1; i != n + 1; i++ {
		if !visited[i] {
			cc++
			ans = append(ans, dfs(visited, g, i, make([]int, 0)))
		}
	}

	fmt.Fprintln(cout, cc)
	for _, comp := range ans {
		fmt.Fprintf(cout, "%d\n%v\n", len(comp), strings.Trim(fmt.Sprint(comp), "[]"))
	}

	cout.Flush()
}

func dfs(visited []bool, g [][]int, x int, ans []int) []int {
	visited[x] = true
	ans = append(ans, x)
	for _, v := range g[x] {
		if !visited[v] {
			ans = dfs(visited, g, v, ans)
		}
	}
	return ans
}