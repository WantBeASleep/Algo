package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, k int
	fmt.Fscan(cin, &n, &k)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	hf := hint{}
	hf.nums = make(map[int]int)
	
	ans := 0
	for _, num := range arr {
		hf.add(num)
		for hf.cntuniq > k {
			hf.remove(arr[hf.l])
		}

		ans += hf.r - hf.l
	}

	fmt.Fprintln(cout, ans)

	cout.Flush()
}

type hint struct {
	l, r int
	nums map[int]int
	cntuniq int
}

func (hf *hint) add(x int) {
	hf.r++
	hf.nums[x]++
	if hf.nums[x] == 1 {
		hf.cntuniq++
	}
}

func (hf *hint) remove(x int) {
	hf.l++
	hf.nums[x]--
	if hf.nums[x] == 0 {
		hf.cntuniq--
	}
}
