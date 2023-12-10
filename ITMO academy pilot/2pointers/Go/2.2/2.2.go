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

	sg := seg{}
	ans := n + 1
	for _, num := range arr {
		sg.add(num)
		for (sg.sum - arr[sg.l]) >= s {
			sg.remove(arr[sg.l])
		}

		if (sg.sum >= s) && (sg.r - sg.l) < ans {
			ans = sg.r - sg.l
		}
	}

	if ans == n + 1 {
		fmt.Fprintln(cout, -1)
	} else {
		fmt.Fprintln(cout, ans)
	}

	cout.Flush()
}

type seg struct {
	l, r int
	sum int
}

func (sg *seg) add(x int) {
	sg.r++
	sg.sum += x
}

func (sg *seg) remove(x int) {
	sg.l++
	sg.sum -= x
}