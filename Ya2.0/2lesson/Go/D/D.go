package main

import "fmt"

func main () {
	var len, k int
	fmt.Scan(&len, &k)
	var a []int = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&a[i])
	}

	lavka(a, len)
}

func iscentral(len int, pos int) bool {
	if len % 2 == 0 {
		return false
	}
	if pos == len / 2 {
		return true
	}
	return false
}

func lavka(a []int, lenL int) {
	var left, right, central int = -1, -1, -1
	for i := 0; i < len(a); i++ {
		if a[i] < lenL / 2 {
			left = i
		}
		if iscentral(lenL, a[i]) {
			central = i
		}
		if right == -1 && a[i] >= lenL / 2 {
			right = i
		}
	}

	if central != -1 {
		fmt.Println(a[central])
	} else {
		fmt.Println(a[left], a[right])
	}
}