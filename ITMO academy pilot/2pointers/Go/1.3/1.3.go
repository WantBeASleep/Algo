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
	arr1 := make([]int, n)
	arr2 := make([]int, m)

	for i := range arr1 {
		fmt.Fscan(cin, &arr1[i])
	}

	for i := range arr2 {
		fmt.Fscan(cin, &arr2[i])
	}

	var ans int
	var i, num, numcnt int
	num = int(-1e9 - 1)
	for j := range arr2 {
		for i != n && arr1[i] <= arr2[j] {
			if num != arr1[i] {
				num = arr1[i]
				numcnt = 1
			} else {
				numcnt++
			}
			i++
		}

		if arr2[j] == num {
			ans += numcnt
		}
	}

	fmt.Fprintln(cout, ans)
	cout.Flush()
}