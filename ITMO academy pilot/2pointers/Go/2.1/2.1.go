package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, s int
	fmt.Fscan(cin, &n, &s)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	sg := segment{limit: s}
	for _, num := range arr {
		sg.add(num)
		for !sg.good() {
			sg.remove(arr[sg.l])
		}
	}

	fmt.Fprintln(cout, sg.maxlen)

	cout.Flush()
}

type segment struct {
	l, r int
	clen, maxlen int
	sum, limit int
}

func (sg *segment) add(x int) {
	sg.r++
	sg.clen++
	sg.sum += x

	if sg.good() && sg.clen > sg.maxlen {
		sg.maxlen = sg.clen
	}
}

func (sg *segment) remove(x int) {
	sg.l++
	sg.clen--
	sg.sum -= x
}

func (sg *segment) good() bool {
	return sg.sum <= sg.limit
}