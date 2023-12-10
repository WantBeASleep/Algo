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
	fmt.Fscan(cin, &n, &m)

	table := make([][]int, n+1)
	for i := 1; i != n+1; i++ {
		table[i] = make([]int, m+1)
		for j := 1; j != m+1; j++ {
			fmt.Fscan(cin, &table[i][j])
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i != n+1; i++ {
		for j := 1; j != m+1; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + table[i][j]
		}
	}

	ans := make([]rune, 0, (n-1)+(m-1))
	i, j := n, m
	for !(i == 1 && j == 1) {
		prevDp := dp[i][j] - table[i][j]
		if dp[i-1][j] == prevDp {
			ans = append(ans, 'D')
			i = i - 1
		} else {
			ans = append(ans, 'R')
			j = j - 1
		}
	}

	fmt.Fprintln(cout, dp[n][m])

	for i := len(ans) - 1; i != -1; i-- {
		fmt.Fprintf(cout, "%c ", ans[i])
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
