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
	var m, n int
	fmt.Fscan(cin, &m, &n)
	stats := make([][3]int, n)

	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &stats[i][0], &stats[i][1], &stats[i][2])
	}

	res := bin(m, n, stats)
	fmt.Fprintln(cout, res)

	getDistrub(m, n, res, stats)
	
	cout.Flush()
}

func good(m, n, ti int, stats [][3]int) bool {
	var ans int

	for i := 0; i != n; i++ {
		t := stats[i][0]
		z := stats[i][1]
		y := stats[i][2]
		
		fullLoop := t * z + y

		ballons := (ti / fullLoop) * z
		mod := (ti % fullLoop) / t
		ballons += int(math.Min(float64(z), float64(mod)))

		ans += ballons
	}

	return ans >= m
}

func bin(m, n int, stats[][3]int) int {
	l := -1
	r := 15000 * 200 + 1

	for r != l + 1 {
		mid := (l + r) / 2
		if good(m, n, mid, stats) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}

func getDistrub(m, n, ti int, stats[][3]int) {
	var ans int

	for i := 0; i != n; i++ {
		t := stats[i][0]
		z := stats[i][1]
		y := stats[i][2]
		
		fullLoop := t * z + y

		ballons := (ti / fullLoop) * z
		mod := (ti % fullLoop) / t
		ballons += int(math.Min(float64(z), float64(mod)))

		if ans + ballons <= m {
			fmt.Fprintf(cout, "%d ", ballons)
		} else {
			if ans <= m {
				fmt.Fprintf(cout, "%d ", m - ans)
			} else {
				fmt.Fprintf(cout, "%d ", 0)
			}
		}

		ans += ballons
	}
}