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
	var n int
	fmt.Fscan(cin, &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(cin, &arr[i])
	}

	hf := hint{}
	hf.init()
	ans := n + 1
	for _, num := range arr {
		hf.add(num)
		for hf.nod() == 1 {
			ans = int(math.Min(float64(ans), float64(hf.r - hf.l)))
			hf.remove()
		}
	}

	if ans == n + 1 {
		fmt.Fprintln(cout, -1)
	} else {
		fmt.Fprintln(cout, ans)
	}
	cout.Flush()
}

type hint struct {
	l, r int
	rstack, lstack stack
}

func (hf *hint) nod() int {
	if hf.lstack.len() == 0 && hf.rstack.len() == 0 {
		return -1
	}
	if hf.lstack.len() != 0 && hf.rstack.len() != 0 {
		return evklid(hf.lstack.nod(), hf.rstack.nod())
	} else if hf.lstack.len() != 0 {
		return hf.lstack.nod()
	} else {
		return hf.rstack.nod()
	}
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

type stack struct {
	data []int
	dnod []int
}

func (st *stack) init() {
	st.data = make([]int, 0)
	st.dnod = make([]int, 0)
}

func (st *stack) len() int {
	return len(st.data)
}

func (st *stack) add(x int) {	
	st.data = append(st.data, x)

	if len(st.dnod) == 0 {
		st.dnod = append(st.dnod, x)
	} else {
		st.dnod = append(st.dnod, evklid(st.dnod[len(st.dnod) - 1], x))
	}
}

func (st *stack) remove() int {
	del := st.data[len(st.data) - 1]
	st.data = st.data[:len(st.data) - 1]
	st.dnod = st.dnod[:len(st.dnod) - 1]
	return del
}

func (st *stack) nod() int {
	return st.dnod[len(st.dnod) - 1]
}

func evklid(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return a + b
}