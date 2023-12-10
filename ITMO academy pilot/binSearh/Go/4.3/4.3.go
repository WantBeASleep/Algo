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

	var n, k int
	fmt.Fscan(cin, &n, &k)

	arr := make([]pair, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &arr[i].a, &arr[i].b)
	}

	fmt.Fprintln(cout, bin(arr, k))

	cout.Flush()
}

type pair struct {
	a, b float64
}

func good(arr []pair, k int, T float64) bool {
	newArr := make([]float64, len(arr))
	for i, e := range arr {
		newArr[i] = e.a - T * e.b
	}

	sort.Slice(newArr, func(i, j int) bool {return newArr[i] > newArr[j]})

	sum := 0.0
	for i := 0; i != k; i++ {
		sum += newArr[i]
	}

	return sum >= 0
}

func bin(arr []pair, k int) float64 {
	l := 0.0
	r := 1.0
	for good(arr, k, r) {
		r *= 2
	}

	for i := 0; i != 100; i++ {
		mid := (l + r) / 2
		if good(arr, k, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}