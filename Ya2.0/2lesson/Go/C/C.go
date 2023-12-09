package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	fmt.Println(countdiff(a))
}

func countdiff(a string) int {
	var count int
	for i := 0; i < len(a) / 2; i++ {
		if a[i] != a[len(a) - 1 - i] {
			count++
		}
	}
	return count
}