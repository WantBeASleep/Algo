package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {

	var n, k int
	fmt.Fscan(cin, &k, &n)

	a := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &a[i])
	}

	fmt.Fprintln(cout, bin(k, a))

	cout.Flush()
}

func good(m, k int, a []int) bool {
	var ans int

	for i := 0; i != len(a); i++ {
		ans += int(math.Min(float64(m), float64(a[i])))
	}

	return ans >= m * k
}

func bin(k int, a []int) int {
	l := 0
	r := 1
	for good(r, k, a) {
		r *= 2
	}

	for r != l + 1 {
		mid := (l + r) / 2
		if good(mid, k, a) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}