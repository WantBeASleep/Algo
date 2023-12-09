package main

import (
	"fmt"
	"bufio"
	"os"
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

	var m int
	fmt.Fscan(cin, &m)

	// тоже код говна, тут только писать структуру в бинарке и там уже обрабатывать ненаход

	for i := 0; i != m; i++ {
		var num int
		fmt.Fscan(cin, &num)

		_, left := binSearch(arr, num, func(elem, x int) bool {return elem >= x})
		right, _ := binSearch(arr, num, func(elem, x int) bool {return elem > x})

		if left == n || right == -1 || arr[left] != num {
			fmt.Fprintln(cout, 0, 0)
		} else {
			fmt.Fprintln(cout, left + 1, right + 1)
		}
		
	}
	cout.Flush()
}

func binSearch(arr []int, x int, check func(int, int) bool) (int, int) {
	var l, r int
	l = -1
	r = len(arr)

	for r != l + 1 {
		m := (l + r) / 2
		if check(arr[m], x) {
			r = m
		} else {
			l = m
		}
	}

	return l, r;
}