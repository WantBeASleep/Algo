package main

import (
	"fmt"
	"math"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(getmaxmin(a))
}

func getmaxmin(a [10]int) int {
	var dist [10]int
	var lasts int = -100
	for i := 0; i < 10; i++ {
		if a[i] == 2 {
			lasts = i
		}
		if a[i] == 1 {
			dist[i] = i - lasts
		}
	}
	lasts = 100
	var ans int
	for i := 9; i >= 0; i-- {
		if a[i] == 2 {
			lasts = i
		}
		if a[i] == 1 {
			dist[i] = int(math.Min(float64(dist[i]), float64(lasts - i)))
			ans = int(math.Max(float64(ans), float64(dist[i])))
		}
	}

	return ans
}
