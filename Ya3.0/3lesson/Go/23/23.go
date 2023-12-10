package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type el struct {
	step, ancs int
}

func main() {
	var n int
	fmt.Fscan(cin, &n)

	dp := make([]el, n + 1)
	dp[0] = el{-1, -1}

	for i := 1; i != n + 1; i++ {
		actions := el{dp[i - 1].step, i - 1}

		if i % 2 == 0 {
			if actions.step > dp[i / 2].step {
				actions = el{dp[i / 2].step, i / 2}
			}
		}
		
		if i % 3 == 0 {
			if actions.step > dp[i / 3].step {
				actions = el{dp[i / 3].step, i / 3}
			}
		}

		actions.step++
		dp[i] = actions
	}

	fmt.Fprintln(cout, dp[n].step)
	ans := make([]int, 0, dp[n].step + 1)

	for i := n; i != 0; i = dp[i].ancs {
		ans = append(ans, i)
	}

	for i := len(ans) - 1; i != -1; i-- {
		fmt.Fprintf(cout, "%d ", ans[i])
	}

	fmt.Fprintln(cout)
	cout.Flush()
}