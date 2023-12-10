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

	table := make([][]int, n + 1)
	for i := 1; i != n + 1; i++ {
		table[i] = make([]int, m + 1)
		for j := 1; j != m + 1; j++ {
			fmt.Fscan(cin, &table[i][j])
		}
	}

	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, m + 1)
	}

	for i := 2; i != m + 1; i++ {
		dp[0][i] = n * m * 100 + 1
	}

	for i := 2; i != n + 1; i++ {
		dp[i][0] = n * m * 100 + 1
	}

	for i := 1; i != n + 1; i++ {
		for j := 1; j != m + 1; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + table[i][j]
		}
	}

	fmt.Fprintln(cout, dp[n][m])

	cout.Flush()
}