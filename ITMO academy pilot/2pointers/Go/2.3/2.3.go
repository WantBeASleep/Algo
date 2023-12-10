package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, s int
	fmt.Fscan(cin, &n, &s)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	hf := hint{}
	ans := 0
	for _, num := range arr {
		hf.add(num)
		for !good(hf.sum, s) {
			hf.remove(arr[hf.l])
		}

		ans += hf.r - hf.l
	}

	fmt.Fprintln(cout, ans)
	cout.Flush()
}

type hint struct {
	l, r int
	sum int
}

func (hi *hint) add(x int) {
	hi.r++
	hi.sum += x
}

func (hi *hint) remove(x int) {
	hi.l++
	hi.sum -= x
}

func good(sum, s int) bool {
	return sum <= s
}