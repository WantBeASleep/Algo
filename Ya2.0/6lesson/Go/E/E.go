package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var n, k int
	fmt.Fscan(cin, &n, &k)	

	arr := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	sort.Ints(arr)

	l := -1
	var r int = 1e18 

	for r != l + 1 {
		m := (l + r) / 2
		if check(arr, k, m) {
			r = m
		} else {
			l = m
		}
	}
	
	fmt.Fprintln(cout, r)
	cout.Flush()
}

func check(arr []int, k, l int) bool {
	cnt := 0
	var lasticld int = -1e9 - 1
	for i := 0; i != len(arr); i++ {
		if lasticld < arr[i] {
			cnt++
			lasticld = arr[i] + l
		}
		if cnt > k {
			return false
		}
	}
	return true
}