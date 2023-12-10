package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscan(cin, &n)

	if n == 0 {
		fmt.Fprintf(cout, "%d\n%d %d\n", 0, 0, 0)
		cout.Flush()
		return
	}

	INF := n*300 + 1
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := range dp[i] {
			if j > i {
				dp[i][j] = INF
			}
		}
	}

	price := make([]int, n+1)
	for i := 1; i != n+1; i++ {
		fmt.Fscan(cin, &price[i])
	}

	dp[1][0] = price[1]
	if price[1] > 100 {
		dp[1][1] = price[1]
	} else {
		dp[1][1] = INF
	}

	for i := 2; i != n+1; i++ {
		for j := 0; j != n+1; j++ {
			if dp[i][j] == INF {
				continue
			}

			free := dp[i-1][j+1]
			money := dp[i-1][j]
			if j > 0 && price[i] > 100 {
				money = dp[i-1][j-1]
			}

			dp[i][j] = min(money+price[i], free)
		}
	}

	totalPrice := INF
	var k1, k2 int
	kdays := make([]int, 0)
	for i := range dp[n] {
		if totalPrice >= dp[n][i] {
			totalPrice = dp[n][i]
			k1 = i
		}
	}

	// for i := range dp {
	// 	fmt.Fprintln(cout, dp[i])
	// }
	// cout.Flush()

	var j = k1
	for i := n; i != 0; i-- {
		if dp[i][j] == dp[i-1][j+1] {
			k2++
			kdays = append(kdays, i)
			j++
		} else {
			if j > 0 && price[i] > 100 {
				j--
			}
		}
	}

	fmt.Fprintln(cout, totalPrice)
	fmt.Fprintln(cout, k1, k2)
	for i := len(kdays) - 1; i != -1; i-- {
		fmt.Fprintf(cout, "%d ", kdays[i])
	}

	fmt.Fprintln(cout)
	cout.Flush()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
