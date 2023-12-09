package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var n, q int

	fmt.Fscan(cin, &n, &q)

	data := make([]int64, n)

	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &data[i])
	}

	suff := make([]int64, n + 1)
	for i := 1; i != n + 1; i++ {
		suff[i] = suff[i - 1] + data[i - 1]
	}

	for i := 0; i != q; i++ {
		var l, r int
		fmt.Fscan(cin, &l, &r)
		l--
		fmt.Fprintln(cout, suff[r] - suff[l])
	}

	cout.Flush()
}