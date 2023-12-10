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

	var i, j int
	for i != n || j != m {
		if j == m || i != n && arr1[i] <= arr2[j] {
			fmt.Fprintf(cout, "%d ", arr1[i])
			i++
		} else {
			fmt.Fprintf(cout, "%d ", arr2[j])
			j++
		}
	}

	fmt.Fprintln(cout)
	cout.Flush()
}