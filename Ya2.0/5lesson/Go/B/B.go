package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

/*
https://github.com/GAlekseyV/YandexTraining/blob/main/src/HW_B5/Task_B/main.cpp - тоже хорошее решение на префф сумах, только написано как говно
*/

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(cin, &n)

	a := make([]int, n)

	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &a[i])
	}

	var maxsum, cursum int

	cursum = a[0]
	maxsum = cursum

	for i := 1; i != n; i++ {
		if a[i] >= 0 {
			if cursum <= 0 {
				cursum = 0
			}
		} else {
			maxsum = int(math.Max(float64(maxsum), float64(cursum)))
		}
		cursum += a[i]
	}

	maxsum = int(math.Max(float64(maxsum), float64(cursum)))

	fmt.Fprintln(cout, maxsum)

	cout.Flush()
}