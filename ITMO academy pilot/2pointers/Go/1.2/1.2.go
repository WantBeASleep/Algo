package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, m int
	fmt.Fscan(cin, &n, &m)
	arr1 := make([]int, n)
	arr2 := make([]int, m)

	for i := range arr1 {
		fmt.Fscan(cin, &arr1[i])
	}

	for i := range arr2 {
		fmt.Fscan(cin, &arr2[i])
	}

	var j int
	for i := range arr2 {
		for j != n && arr1[j] < arr2[i] {
			j++
		}
		fmt.Fprintf(cout, "%d ", j)
	}

	fmt.Fprintln(cout)

	cout.Flush()
}