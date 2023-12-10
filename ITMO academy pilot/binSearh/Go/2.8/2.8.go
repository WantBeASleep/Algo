package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type stats struct {
	b_count, s_count, c_count int
	nb, ns, nc int
	pb, ps, pc int
	r int
}

func main() {

	var str string
	fmt.Fscan(cin, &str)

	var b_count, s_count, c_count int
	for _, char := range str {
		switch char {
		case 'B':
			b_count++
		case 'S':
			s_count++
		case 'C':
			c_count++
		}
	}

	var nb, ns, nc int
	fmt.Fscan(cin, &nb, &ns, &nc)

	var pb, ps, pc int
	fmt.Fscan(cin, &pb, &ps, &pc)

	var r int
	fmt.Fscan(cin, &r)

	data := stats{b_count, s_count, c_count, nb, ns, nc, pb, ps, pc, r}

	fmt.Fprintln(cout, bin(data))

	cout.Flush()
}

func good(data stats, m int) bool {
	var b_buy, s_buy, c_buy int

	if m * data.b_count > data.nb {
		b_buy = m * data.b_count - data.nb
	}
	if m * data.s_count > data.ns {
		s_buy = m * data.s_count - data.ns
	}
	if m * data.c_count > data.nc {
		c_buy = m * data.c_count - data.nc
	}

	return b_buy * data.pb + s_buy * data.ps + c_buy * data.pc <= data.r
}

func bin(data stats) int {
	l := 0
	r := 1
	for good(data, r) {
		r *= 2
	}

	for r != l + 1 {
		mid := (l + r) / 2
		if good(data, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}