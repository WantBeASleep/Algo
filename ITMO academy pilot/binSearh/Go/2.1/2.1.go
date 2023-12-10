package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var w, h, n int
	fmt.Fscan(cin, &w, &h, &n)

	fmt.Fprintln(cout, binSearch(w, h, n))
	
	cout.Flush()
}

func good(x, w, h, n int) bool {
	vert := x / h
	horz := x / w
	return vert * horz >= n
}

func binSearch(w, h, n int) int {
	l := 0
	r := 1
	for !good(r, w, h, n) {
		r *= 2
	}

	for r != l + 1 {
		m := (r + l) / 2
		if good(m, w, h, n) {
			r = m
		} else {
			l = m
		}
	}

	return r
}