package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var s string
	fmt.Fscan(cin, &s)

	z := make([]int, len(s))

	var l, r int
	for i := 1; i != len(z); i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
		}
		for z[i] + i < len(z) && s[i + z[i]] == s[z[i]] {
			z[i]++
		}
		if r < i + z[i] {
			r = i + z[i]
			l = i
		}
	}

	fmt.Fprintln(cout, strings.Trim(fmt.Sprint(z), "[]"))

	cout.Flush()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}