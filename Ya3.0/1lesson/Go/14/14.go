package main

import (
	"fmt"
	"bufio"
	"errors"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscan(cin, &n)

	path1 := make([]int, n)

	for i := range path1 {
		fmt.Fscan(cin, &path1[i])
	}

	path2 := make([]int, 0, n)

	st := stack{stack: make([]int, 0, n)}

	for _, wagon := range path1 {
		top, e := st.back()
		for e == nil && top < wagon {
			path2 = append(path2, top)
			st.pop()
			top, e = st.back()
		}

		st.push(wagon)
	}

	top, e := st.back()
	for e == nil {
		path2 = append(path2, top)
		st.pop()
		top, e = st.back()
	}

	ans := "YES"
	for i, wagon := range path2 {
		if i + 1 != wagon {
			ans = "NO"
		}
	}

	fmt.Fprintln(cout, ans)

	cout.Flush()
}

type stack struct {
	stack []int
}

func (st *stack) push(x int) {
	st.stack = append(st.stack, x)
}

func (st *stack) back() (int, error) {
	if len(st.stack) == 0 {
		return 0, errors.New("stack empty")
	}

	return st.stack[len(st.stack) - 1], nil
}

func (st *stack) pop() {
	st.stack = st.stack[:len(st.stack) - 1]
}