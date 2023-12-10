package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscan(cin, &n)

	h := heap{data: make([]int, n)}
	for i := 0; i != n; i++ {
		fmt.Fscan(cin, &h.data[i])
	}

	
	for i := (n - 2) / 2; i != -1; i-- {
		h.downSort(i)
	}

	for i := 0; i != n; i++ {
		fmt.Fprintf(cout, "%d ", h.extract())
	}


	fmt.Fprintln(cout)
	cout.Flush()
}

type heap struct {
	data []int
}

func haveChilds(x, lim int) bool {
	return 2 * x + 1 < lim
}

func (h *heap) lesschild(x int) bool {
	if 2 * x + 2 < len(h.data) {
		return h.data[x] > h.data[2 * x + 1] || h.data[x] > h.data[2 * x + 2]
	}
	return h.data[x] > h.data[2 * x + 1]
}

func (h *heap) downSort(_idx int) {
	idx := _idx
	for haveChilds(idx, len(h.data)) && h.lesschild(idx) {
		var switchIdx int
		
		if 2 * idx + 2 < len(h.data) {
			if h.data[2 * idx + 1] < h.data[2 * idx + 2] {
				switchIdx = 2 * idx + 1
			} else {
				switchIdx = 2 * idx + 2
			}
		} else {
			switchIdx = 2 * idx + 1
		}

		h.data[idx], h.data[switchIdx] = h.data[switchIdx], h.data[idx]
		idx = switchIdx
	}
}

func (h *heap) extract() (ans int) {
	ans = h.data[0]

	h.data[0] = h.data[len(h.data) - 1]
	h.downSort(0)

	h.data = h.data[:len(h.data) - 1]
	
	return
}