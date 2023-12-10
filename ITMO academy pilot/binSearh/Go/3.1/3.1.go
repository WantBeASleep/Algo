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
	var n int
	fmt.Fscan(cin, &n)

	st := make([][2]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &st[i][0], &st[i][1])
	}

	fmt.Fprintln(cout, bin(st))

	cout.Flush()
}

type segment struct {
	left float64
	right float64
	statusZero bool
}

func (sg *segment) cross(left, right float64) bool {
	if !sg.statusZero {
		sg.left = left
		sg.right = right
		sg.statusZero = true
	} else {
		sg.left = math.Max(sg.left, left)
		sg.right = math.Min(sg.right, right)
	}

	return sg.left <= sg.right
}

func good(ti float64, st [][2]int) bool {
	var sg segment
	for i := 0; i != len(st); i++ {
		left := float64(st[i][0]) - ti * float64(st[i][1])
		right := float64(st[i][0]) + ti * float64(st[i][1])

		flagNonEmpty := sg.cross(left, right)
		if !flagNonEmpty {
			return false
		}
	}
	return true
}

func bin(st[][2]int) float64{
	l := 0.0
	r := 1.0
	for !good(r, st) {
		r *= 2
	}

	for i := 0; i != 100; i++ {
		mid := (l + r) / 2
		if good(mid, st) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}