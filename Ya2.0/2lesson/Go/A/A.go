package main

import "fmt"

func main() {
	var count int
	count = maxelemcount()
	fmt.Println(count)
}

func maxelemcount() int {
	var curVal, maxVal, count int
	fmt.Scan(&curVal);
	count = 0
	maxVal = curVal;

	for curVal != 0 {
		if (curVal > maxVal) {
			maxVal = curVal
			count = 1
		} else if (curVal == maxVal) {
			count++
		}
		fmt.Scan(&curVal)
	}

	return count
}