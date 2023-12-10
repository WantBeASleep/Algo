package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type st struct {
	A, B, C int
}

func main() {
	var n int
	fmt.Fscan(cin, &n)
	stat := make([]st, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &stat[i].A, &stat[i].B, &stat[i].C)
	}

	dp := make([]int, n)

	dp[0] = stat[0].A
	if n >= 2 {
		dp[1] = min(dp[0] + stat[1].A, stat[0].B)
	}
	if n >= 3 {
		dp[2] = min(min(dp[1] + stat[2].A, dp[0] + stat[1].B), stat[0].C)
	}

	for i := 3; i < n; i++ {
		t := dp[i - 1] + stat[i].A
		t = min(t, dp[i - 2] + stat[i - 1].B)
		t = min(t, dp[i - 3] + stat[i - 2].C)

		dp[i] = t
	}

	fmt.Fprintln(cout, dp[n - 1])
	cout.Flush()
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}