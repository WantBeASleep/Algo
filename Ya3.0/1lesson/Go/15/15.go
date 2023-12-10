package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type town struct {
	idx, price int
}

func main() {
	var n int
	fmt.Fscan(cin, &n)
	prices := make([]town, n)
	ans := make([]int, n)
	st := stack{stack: make([]town, 0, n)}

	for i := range prices{
		var price int
		fmt.Fscan(cin, &price)

		prices[i] = town{idx: i, price: price}
	}

	for _, city := range prices {
		top, e := st.back()
		for e == nil && top.price > city.price {
			ans[top.idx] = city.idx
			st.pop()
			top, e = st.back()
		}

		st.push(city)
	}

	top, e := st.back()
	for e == nil {
		ans[top.idx] = -1
		st.pop()
		top, e = st.back()
	}

	fmt.Fprintln(cout, strings.Trim(fmt.Sprint(ans), "[]"))
	

	cout.Flush()
}

type stack struct {
	stack []town
}

func (st *stack) push(x town) {
	st.stack = append(st.stack, x)
}

func (st *stack) back() (town, error) {
	if len(st.stack) == 0 {
		return town{}, errors.New("stack empty")
	}

	return st.stack[len(st.stack) - 1], nil
}

func (st *stack) pop() {
	st.stack = st.stack[:len(st.stack) - 1]
}