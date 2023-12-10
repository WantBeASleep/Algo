package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cin = bufio.NewScanner(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var N int
	fmt.Scan(&N)

	hipar := heap{data: make([]int, 0, 100)}

	for i := 0; i != N; i++ {
		cin.Scan()
		line := strings.Split(cin.Text(), " ")

		typeOp, _ := strconv.Atoi(line[0])
		if typeOp == 0 {
			num, _ := strconv.Atoi(line[1])
			hipar.insert(num)
		} else {
			fmt.Fprintln(cout, hipar.extract())
		}
	}

	cout.Flush()
}

type heap struct {
	data []int
}

func (h *heap) insert(x int) {
	h.data = append(h.data, x)

	// просеивание вверх юзается только тут, так что фунции не будет
	idx := len(h.data) - 1
	for idx != 0 && h.data[(idx - 1) / 2] < h.data[idx] {
		h.data[idx], h.data[(idx - 1) / 2] = h.data[(idx - 1) / 2], h.data[idx] // swap
		idx = (idx - 1) / 2
	}
}

func (h *heap) extract() (ans int) {
	ans = h.data[0]

	h.data[0] = h.data[len(h.data) - 1]
	idx := 0
	for (2 * idx + 2 < len(h.data)) && (h.data[idx] < h.data[2 * idx + 1] || h.data[idx] < h.data[2 * idx + 2]) {
		var switchIdx int
		left := h.data[2 * idx + 1]
		right := h.data[2 * idx + 2]
		
		if left < right {
			switchIdx = 2 * idx + 2
		} else {
			switchIdx = 2 * idx + 1
		}

		h.data[idx], h.data[switchIdx] = h.data[switchIdx], h.data[idx]
		idx = switchIdx
	}

	h.data = h.data[:len(h.data) - 1]
	
	return
}