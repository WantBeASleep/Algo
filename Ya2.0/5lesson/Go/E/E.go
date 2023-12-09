package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elem struct {
	el int
	idx int
}

type ans struct {
	i, j, k int
	real bool
}

func comp(a, b ans) bool {
	if a.i == b.i {
		if a.j == b.j {
			return a.k < b.k
		} else {
			return a.j < b.j
		}
	} else {
		return a.i < b.i
	}
}

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)

	var s int
	fmt.Fscan(cin, &s)

	var sz1, sz2, sz3 int

	fmt.Fscan(cin, &sz1)
	A := make([]elem, sz1)
	amap := make(map[int]int)
	var last int
	var ireal int
	for i := 0; i != sz1; i++ {
		var num int
		fmt.Fscan(cin, &num)
		_, flag := amap[num]

		if !flag {
			A[ireal] = elem{el: num, idx: ireal}
			last = ireal
			amap[num] = 1
			ireal++
		}
	}
	A = A[:last + 1]

	fmt.Fscan(cin, &sz2)
	B := make([]elem, sz2)
	bmap := make(map[int]int)
	ireal = 0
	for i := 0; i != sz2; i++ {
		var num int
		fmt.Fscan(cin, &num)
		_, flag := bmap[num]

		if !flag {
			B[ireal] = elem{el: num, idx: ireal}
			last = ireal
			bmap[num] = 1
			ireal++
		}
	}
	B = B[:last + 1]

	fmt.Fscan(cin, &sz3)
	C := make([]elem, sz3)
	cmap := make(map[int]int)
	ireal = 0
	for i := 0; i != sz3; i++ {
		var num int
		fmt.Fscan(cin, &num)
		_, flag := cmap[num]

		if !flag {
			C[ireal] = elem{el: num, idx: ireal}
			last = ireal
			cmap[num] = 1
			ireal++
		}
	}
	C = C[:last + 1]

	sort.Slice(A, func(i, j int) bool {return A[i].el < A[j].el})
	sort.Slice(B, func(i, j int) bool {return B[i].el < B[j].el})
	sort.Slice(C, func(i, j int) bool {return C[i].el < C[j].el})

	answr := ans{real: false}
	for i := 0; i != len(A); i++ {
		var k = len(C) - 1
		for j := 0; j != len(B); j++ {
			for k >= 0 && A[i].el + B[j].el + C[k].el > s {
				k--
			}
			if k < 0 {
				break 
			}

			if A[i].el + B[j].el + C[k].el == s {
				if !answr.real || comp(ans{A[i].idx, B[j].idx, C[k].idx, true}, answr) {
					answr = ans{A[i].idx, B[j].idx, C[k].idx, true}
				}	
			}
		}
	}

	if !answr.real {
		fmt.Fprintln(cout, -1)
	} else {
		fmt.Fprintln(cout, answr.i, answr.j, answr.k)
	}
	
	cout.Flush()
}