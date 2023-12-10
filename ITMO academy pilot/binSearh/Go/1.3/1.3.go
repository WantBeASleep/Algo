package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var N, M int
	fmt.Fscan(cin, &N, &M)

	arr := make([]int, N)
	for i := 0; i != N; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	for i := 0; i != M; i++ {
		var num int
		fmt.Fscan(cin, &num)

		fmt.Fprintln(cout, binSearh(arr, num))
	}
	
	cout.Flush()
}

func binSearh(arr []int, x int) int {
	l := -1
	r := len(arr)

	for r != l + 1 {
		m := (l + r) / 2
		if arr[m] >= x {
			r = m
		} else {
			l = m
		}
	}

	return r + 1
}