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
	
	var a, b, c, d float64
	fmt.Fscan(cin, &a, &b, &c, &d)

	fmt.Fprintf(cout, "%f\n", binFloatSearch(a, b, c, d))
	
	cout.Flush()
}

func binFloatSearch(a, b, c, d float64) float64 {
	var l float64 = -1e10
	var r float64 = 1e10
	var kef bool = true
	if a < 0 {
		kef = false
	}

	for i := 0; i != 1000; i++ {
		m := (l + r) / 2
		if kef == plus(a, b, c, d, m) {
			r = m
		} else {
			l = m
		}
	}

	return l
}

func plus(a, b, c, d, x float64) bool {
	return a * math.Pow(x, 3) + b * math.Pow(x, 2) + c * x + d >= 0
}