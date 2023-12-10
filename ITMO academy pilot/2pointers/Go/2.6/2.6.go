package main

import (
	"bufio"
	"fmt"
	"math"
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
	hf.init()
	ans := 0
	for _, num := range arr {
		hf.add(num)
		for hf.diff() > k {
			hf.remove()
		}

		ans += hf.r - hf.l
	}

	fmt.Fprintln(cout, ans)
	cout.Flush()
}

type hint struct {
	l, r int
	rstack, lstack stack
}

func (hf *hint) init() {
	hf.lstack.init()
	hf.rstack.init()
}

func (hf *hint) add(x int) {
	hf.r++
	hf.rstack.add(x)
}

func (hf *hint) remove() {
	hf.l++
	if hf.lstack.len() == 0 {
		for hf.rstack.len() != 0 {
			hf.lstack.add(hf.rstack.remove())
		}
	}

	hf.lstack.remove()
}

func (hf *hint) diff() int {
	if hf.lstack.len() != 0 && hf.rstack.len() != 0 {
		lmin, lmax := hf.lstack.min(), hf.lstack.max()
		rmin, rmax := hf.rstack.min(), hf.rstack.max()

		return int(math.Max(float64(lmax), float64(rmax))) - int(math.Min(float64(lmin), float64(rmin)))
	} else if hf.lstack.len() != 0 {
		return hf.lstack.max() - hf.lstack.min()
	} else {
		return hf.rstack.max() - hf.rstack.min()
	}
}

type stack struct {
	data []int
	dmin []int
	dmax []int
}

func (st *stack) init() {
	st.data = make([]int, 0)
	st.dmin = make([]int, 0)
	st.dmax = make([]int, 0)
}

func (st *stack) len() int {
	return len(st.data)
}

func (st *stack) add(x int) {
	st.data = append(st.data, x)

	if len(st.dmax) != 0 {
		st.dmax = append(st.dmax, int(math.Max(float64(x), float64(st.dmax[len(st.dmax) - 1]))))
	} else {
		st.dmax = append(st.dmax, x)
	}
	
	if len(st.dmin) != 0 {
		st.dmin = append(st.dmin, int(math.Min(float64(x), float64(st.dmin[len(st.dmin) - 1]))))
	} else {
		st.dmin = append(st.dmin, x)
	}
}

func (st *stack) remove() int {
	del := st.data[len(st.data) - 1]
	st.data = st.data[:len(st.data) - 1]
	st.dmax = st.dmax[:len(st.dmax) - 1]
	st.dmin = st.dmin[:len(st.dmin) - 1]
	return del
}

func (st *stack) min() int {
	return st.dmin[st.len() - 1]
}

func (st *stack) max() int {
	return st.dmax[st.len() - 1]
}