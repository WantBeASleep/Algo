package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, p int
	fmt.Fscan(cin, &n, &p)

	arr := make([]int, n)
	sumOfArr := 0
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
		sumOfArr += arr[i]
	}

	arr = append(arr, arr...)

	between := p / sumOfArr
	lenSg := len(arr)
	sumSg := 0
	l := 0
	r := 0
	left := 0 

	if between != 0 {
		r = len(arr) / 2
		sumSg = (between + 1) * sumOfArr
		for sumSg - arr[l] >= p && l != r{
			sumSg -= arr[l]
			l++
		}

		lenSg = r + (between * n) - l
		left = l % (len(arr) / 2)
	}

	for r != len(arr) && l != len(arr) / 2 {
		sumSg += arr[r]
		r++

		for sumSg - arr[l] >= p && l != r && l != len(arr) / 2 {
			sumSg -= arr[l]
			l++
		}

		if sumSg >= p {
			if r + (between * n) - l < lenSg {
				lenSg = r + (between * n) - l
				left = l
			}
		}
	}

	fmt.Fprintln(cout, left + 1, lenSg)

	cout.Flush()
}