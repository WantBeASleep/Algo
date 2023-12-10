package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var n, k int
	fmt.Fscan(cin, &n, &k)	

	m := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &m[i])
	}

	for i := 0; i != k; i++ {
		var num int
		fmt.Fscan(cin, &num)

		if binSearch(m, num) {
			fmt.Fprintln(cout, "YES")
		} else {
			fmt.Fprintln(cout, "NO")
		}
	}
	
	cout.Flush()
}

func binSearch(arr []int, x int) bool {
	l := -1
	r := len(arr)

	for r != l + 1 {
		m := (l + r) / 2
		if arr[m] > x {
			r = m
		} else {
			l = m
		}
	}

	if l == -1 || arr[l] != x {
		return false
	} else {
		return true
	}
}