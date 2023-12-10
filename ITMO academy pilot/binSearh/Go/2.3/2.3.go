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
	
	var n, x, y int
	fmt.Fscan(cin, &n, &x, &y)

	fmt.Fprintln(cout, binSearch(x, y, n))
	
	cout.Flush()
}

func good(x, y, n, t int) bool {
	precopy := int(math.Min(float64(x), float64(y)))
	
	if precopy > t {
		return false
	}

	first := (t - precopy) / x
	second := (t - precopy) / y

	return first + second >= n - 1
}

func binSearch(x, y, n int) int {
	l := 0
	r := 1
	for !good(x, y, n, r) {
		r *= 2
	}

	for r != l + 1 {
		m := (l + r) / 2
		if good(x, y, n, m) {
			r = m
		} else {
			l = m
		}
	}

	return r
}