package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var N int
	fmt.Fscan(cin, &N)

	arr := make([]int, N)
	for i := 0; i != N; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	sort.Ints(arr)

	var M int
	fmt.Fscan(cin, &M)

	for i := 0; i != M; i++ {
		var l, r int
		fmt.Fscan(cin, &l, &r)

		left := bin(arr, l)
		right := bin(arr, r + 1)

		fmt.Fprint(cout, right - left, " ")
	}
	
	cout.Flush()
}

func bin(arr []int, x int) int {
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

	return r
}