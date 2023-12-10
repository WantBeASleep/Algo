package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, d int
	fmt.Fscan(cin, &n, &d)

	arr := make([]float64, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	
	l, r := bin(arr, d)
	fmt.Fprintln(cout, l + 1, r) // возвращаем отрезки ебать

	cout.Flush()
}

func good(arr []float64, x float64, d int) (bool, int, int) {
	newArr := make([]float64, len(arr))
	for i, num := range arr {
		newArr[i] = num - x
	}

	preffArr := make([]float64, len(arr) + 1)
	for i := 1; i != len(arr) + 1; i++ {
		preffArr[i] = preffArr[i - 1] + newArr[i - 1]
	}

	minPreffArg := make([]int, len(arr) + 1)
	for i := 1; i != len(arr) + 1; i++ {
		if preffArr[i] < preffArr[minPreffArg[i - 1]] {
			minPreffArg[i] = i
		} else {
			minPreffArg[i] = minPreffArg[i - 1]
		}
	}

	for r := d; r != len(arr) + 1; r++ {
		l := r - d
		if preffArr[r] - preffArr[minPreffArg[l]] >= 0 {
			return true, minPreffArg[l], r
		}
	}

	return false, -1, -1
}

func boolGoodWrapper(arr []float64, x float64, d int) bool {
	ok, _, _ := good(arr, x, d)
	return ok
}

func bin(arr []float64, d int) (int, int) {
	l := 0.0
	r := 1.0
	for boolGoodWrapper(arr, r, d) {
		r *= 2
	}

	for i := 0; i != 150; i++ {
		mid := (l + r) / 2
		if boolGoodWrapper(arr, mid, d) {
			l = mid
		} else {
			r = mid
		}
	}

	_, left, right := good(arr, l, d)
	return left, right
}