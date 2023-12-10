package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type run struct {
	letter rune
	step int
}

func main() {

	var str1, str2 string
	fmt.Fscan(cin, &str1, &str2)

	cross := make([]int, len(str1))
	for i := 0; i != len(str1); i++ {
		fmt.Fscan(cin, &cross[i])
	}
	
	indxStr := make([]run, len(str1))
	for i, num := range cross {
		indxStr[num - 1] = run{letter: rune(str1[num - 1]), step: i + 1}
	}

	fmt.Fprintln(cout, bin(indxStr, str2))

	cout.Flush()
}

func good(indxStr []run, orig string, step int) bool {
	var i, j int

	for i < len(indxStr) && j < len(orig) {
		for i < len(indxStr) && indxStr[i].step <= step {
			i++
		}

		if i == len(indxStr) {
			break
		}

		if indxStr[i].letter == rune(orig[j]) {
			j++
		}
		i++
	}

	return j == len(orig)
}

// ~O(N*LogN)
func bin(indxStr []run, orig string) int {
	l := 0
	r := len(indxStr)

	for r != l + 1 {
		mid := (l + r) / 2
		if good(indxStr, orig, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}