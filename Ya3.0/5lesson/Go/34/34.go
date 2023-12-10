package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)
var answr []int = make([]int, 0)

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
	}

	colorg := make([]int, n + 1)
	haveLoop := false
	for i := 1; i != n + 1; i++ {
		if colorg[i] == 0 {
			haveLoop = haveLoop || dfsL(g, colorg, i)
		}
	}

	if haveLoop {
		fmt.Fprintln(cout, -1)
		cout.Flush()
		return
	}

	
	visited := make([]bool, n + 1)
	for i := 1; i != n + 1; i++ {
		if !visited[i] {
			dfs(g, visited, i)
		}
	}

	for i := range answr {
		fmt.Fprintf(cout, "%d ", answr[len(answr) - 1 - i])
	}

	fmt.Fprintln(cout)

	cout.Flush()
}

// 0 - white
// 1 - grey
// 2 - black
// true - have loop int vertex!!!
func dfsL(g [][]int, colorg []int, x int) bool {
	colorg[x] = 1
	for _, v := range g[x] {
		switch colorg[v]{
		case 0:
			ans := dfsL(g, colorg, v)
			if ans {
				return true
			}
		case 1:
			return true
		}
	}
	colorg[x] = 2
	return false
}

func dfs(g [][]int, visited []bool, x int) {
	visited[x] = true
	for _, v := range g[x] {
		if !visited[v] {
			dfs(g, visited, v)
		}
	}
	answr = append(answr, x)
}