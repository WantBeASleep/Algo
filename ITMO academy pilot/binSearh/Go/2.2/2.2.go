package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var N, K int
	fmt.Fscan(cin, &N, &K)
	arr := make([]float64, N)

	for i := 0; i != N; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	fmt.Fprintln(cout, bin(arr, K))
	
	cout.Flush()
}

func good(arr []float64, k int, ln float64) bool {
	var cnt int
	for _, rope := range arr {
		slices, _ := math.Modf(rope / ln)
		cnt += int(slices)
	}
	return cnt >= k
}

func bin(arr []float64, k int) float64 {
	l := 0.0
	r := 1.0
	for good(arr, k, r) {
		r *= 2
	}

	for i := 0; i != 100; i++ {
		m := (l + r) / 2
		if good(arr, k, m) {
			l = m
		} else {
			r = m
		}
	}

	return l
}