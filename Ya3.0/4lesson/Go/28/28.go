package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

var n, m int

func main() {
	fmt.Fscan(cin, &n, &m)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[1][1] = 1

	for i := 1; i != n+1; i++ {
		for j := 1; j != m+1; j++ {
			if i == 1 && j == 1 {
					continue
			} // Да хуево, но так проще, чем городить забор

			dp[i][j] = at(i-2, j-1, dp) + at(i-1, j-2, dp)
		}
	}

	fmt.Fprintln(cout, dp[n][m])

	cout.Flush()
}

func inTable(i, j int) bool {
	return i >= 1 && i <= n && j >= 1 && j <= m
}

func at(i, j int, dp [][]int) int {
	if !inTable(i, j) {
		return 0
	}
	return dp[i][j]
}
