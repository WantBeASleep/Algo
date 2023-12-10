package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type chains struct {
	k, e, f int
}

func (c chains) cnt() int {
	return c.e + c.f + c.k
}

func main() {
	var n int
	fmt.Fscan(cin, &n)
	dp := make([]chains, n)

	dp[0] = chains{k: 0, e: 1, f: 1}
	if n >= 2 {
		dp[1] = chains{k: 1, e: 2, f: 1}
	}

	for i := 2; i < n; i++ {
		dp[i].k = dp[i - 1].f
		dp[i].e = dp[i - 1].k + dp[i - 1].e + dp[i - 1].f
		dp[i].f = dp[i - 1].e
	}

	fmt.Fprintln(cout, dp[n - 1].cnt())

	cout.Flush()
}

