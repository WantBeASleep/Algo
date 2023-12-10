package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cin = bufio.NewReaderSize(os.Stdin, 1000000)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	line, _ := cin.ReadString('\n')
	newLine := line[:len(line) - 1]

	argLine := strings.Split(newLine, " ")
	st := stack{stack: make([]int, 0)}

	for _, elem := range argLine {
		if elem == "" {
			continue
		}
		num, e := strconv.Atoi(elem)

		if e == nil {
			st.push(num)
		} else {
			arg2, _ := st.back()
			st.pop()
			arg1, _ := st.back()
			st.pop()

			switch elem{
			case "+":
				st.push(arg1 + arg2)
			case "-":
				st.push(arg1 - arg2)
			case "*":
				st.push(arg1 * arg2)
			}
		}
	}


	ans, _ := st.back()
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