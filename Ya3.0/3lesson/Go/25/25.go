package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscan(cin, &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	sort.Ints(arr)

	dp := make([]int, n)
	dp[1] = arr[1] - arr[0]

	if n >= 3 {
		dp[2] = dp[1] + arr[2] - arr[1]
	}

	for i := 3; i < n; i++ {
		diff := arr[i] - arr[i - 1]
		dp[i] = min(dp[i - 2] + diff, dp[i - 1] + diff)
	}

	fmt.Fprintln(cout, dp[n - 1])

	cout.Flush()
}

