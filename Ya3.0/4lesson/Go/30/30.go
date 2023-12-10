package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, m int

	fmt.Fscan(cin, &n)
	first := make([]int, n+1)
	for i := 1; i != n+1; i++ {
		fmt.Fscan(cin, &first[i])
	}

	fmt.Fscan(cin, &m)
	second := make([]int, m+1)
	for i := 1; i != m+1; i++ {
		fmt.Fscan(cin, &second[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// first[0], second[0] = 10001, -10001

	for i := 1; i != n+1; i++ {
		for j := 1; j != m+1; j++ {
			if first[i] == second[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	ans := make([]int, 0)
	var i, j = n, m
	for i > 0 && j > 0 {
		switch dp[i][j] {
		case dp[i-1][j]:
			i--
		case dp[i][j-1]:
			j--
		default:
			ans = append(ans, first[i])
			i--
			j--
		}
	}

	for i := len(ans) - 1; i != -1; i-- {
		fmt.Fprintf(cout, "%d ", ans[i])
	}
	fmt.Fprintln(cout)

	cout.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
