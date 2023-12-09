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
	
	var a, k, b, m, x int
	fmt.Fscan(cin, &a, &k, &b, &m, &x)

	fmt.Fprintln(cout, binSearch(a, k, b, m, x))
	
	cout.Flush()
}

func binSearch(a, k, b, m, x int) int {
	l := 0
	r := 2 * x / int(math.Max(float64(a), float64(b))) + 1

	for r != l + 1 {
		mid := (l + r) / 2
		if check(a, k, b, m, x, mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func check(a, k, b, m, x, day int) bool {
	workA := day - (day / k)
	workB := day - (day / m)
	return (a * workA + b * workB) >= x
}