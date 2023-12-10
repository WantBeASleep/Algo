package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, k int
	fmt.Fscan(cin, &n, &k)

	arr := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	fmt.Fprintln(cout, bin(arr, k))

	cout.Flush()
}

func good(arr []int, k, t int) bool {
	cowInStall := 0
	lastStall := -t // шобы первое стойло сработало

	for i := 0; i != len(arr); i++ {
		if arr[i] - lastStall >= t {
			cowInStall++
			lastStall = arr[i]
		}
	}

	return cowInStall >= k
}

func bin(arr []int, k int) int {
	l := 0
	r := 1
	for good(arr, k, r) {
		r *= 2
	}

	for r != l + 1 {
		mid := (l + r) / 2
		if good(arr, k, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
