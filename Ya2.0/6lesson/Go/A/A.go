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
	
	var n int
	fmt.Fscan(cin, &n)

	arr := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &arr[i])
	}

	sort.Ints(arr)

	var m int
	fmt.Fscan(cin, &m)

	for i := 0; i != m; i++ {
		var l, r int
		fmt.Fscan(cin, &l, &r)

		// бинарка вернет max i : a_i < x
		// работаем на полуинтервалах
		left := binSearch(arr, l) + 1
		right := binSearch(arr, r + 1) + 1

		fmt.Fprint(cout, right - left, " ")
	}
	
	fmt.Fprintln(cout)
	cout.Flush()
}


func binSearch(a []int, x int) int {
	l := -1
	r := len(a)
	for r != l + 1 {
		m := (l + r) / 2
		if lowerBinCheck(a[m], x) {
			r = m
		} else {
			l = m
		}
	}
	return l;
}

func lowerBinCheck(elem, l int) bool {
	return elem >= l 
}
