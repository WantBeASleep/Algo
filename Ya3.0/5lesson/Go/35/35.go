package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

// EMPTY, FILLING, DONE
type loopV struct {
	loop []int
	status string
}

var ans loopV = loopV{make([]int, 0), "EMPTY"}

func main() {
	var n int
	fmt.Fscan(cin, &n)
	g := make([][]int, n + 1)
	for i := 1; i != n + 1; i++ {
		g[i] = make([]int, 0)
		for j := 1; j != n + 1; j++ {
			var edge int
			fmt.Fscan(cin, &edge)

			if edge == 1 {
				g[i] = append(g[i], j)
			}
		}
	}

// Боже как же плохо это написано, но ладно, это рекурсия, сойдет

	visited := make([]bool, n + 1)
	for i := 1; i != n + 1; i++ {
		if !visited[i] {
			res := dfsL(g, visited, i, -1)

			if res {
				PrintZalupa()
				cout.Flush()
				return
			}
		}
	}

	fmt.Fprintln(cout, "NO")

	cout.Flush()
}

func PrintZalupa() {
	fmt.Fprintln(cout, "YES")
	fmt.Fprintln(cout, len(ans.loop))
	fmt.Fprintln(cout, strings.Trim(fmt.Sprint(ans.loop), "[]"))
}

// y - ancs
func dfsL(g [][]int, visited []bool, x, y int) bool {
	visited[x] = true
	for _, v := range g[x] {
		if v == y {
			continue
		}
		if visited[v] {
			ans.loop = append(ans.loop, v, x)
			ans.status = "FILLING"
			return true
		}
		if dfsL(g, visited, v, x) {
			if ans.status == "DONE" {
				return true
			}

			if x == ans.loop[0] {
				ans.status = "DONE"
			} else {
				ans.loop = append(ans.loop, x)
			}
			return true
		}
	}
	return false
}