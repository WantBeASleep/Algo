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
	var c float64
	fmt.Fscan(cin, &c)

	fmt.Fprintln(cout, bin(c))

	cout.Flush()
}

func good(c, x float64) bool {
	return math.Pow(x, 2) + math.Sqrt(x) > c
}

func bin(c float64) float64 {
	l := 0.0
	r := 1.0
	for !good(c, r) {
		r *= 2
	}

	for i := 0; i != 150; i++ {
		mid := (l + r) / 2
		if good(c, mid) {
			r = mid
		} else {
			l = mid
		}
	}
	
	return r
}