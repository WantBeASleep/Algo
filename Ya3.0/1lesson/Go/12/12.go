package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var line string
	errLine := false
	fmt.Fscan(cin, &line)

	st := stack{stack: make([]rune, 0)}
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' {
			st.push(char)
		} else {
			switch char {
			case ')':
				last, e := st.back()
				if e != nil || last != '(' {
					errLine = true
					break
				} else {
					st.pop()
				}
			case ']':
				last, e := st.back()
				if e != nil || last != '[' {
					errLine = true
					break
				} else {
					st.pop()
				}
			case '}':
				last, e := st.back()
				if e != nil || last != '{' {
					errLine = true
					break
				} else {
					st.pop()
				}
			}
		}
	}

	if !errLine && st.size() == 0 {
		fmt.Fprintln(cout, "yes")
	} else {
		fmt.Fprintln(cout, "no")
	}

	cout.Flush()
}

type stack struct {
	stack []rune
}

func (st *stack) size() int {
	return len(st.stack)
}

func (st *stack) push(x rune) {
	st.stack = append(st.stack, x)
}

func (st *stack) back() (rune, error) {
	if len(st.stack) == 0 {
		return 0, errors.New("stack empty")
	}

	return st.stack[len(st.stack) - 1], nil
}

func (st *stack) pop() {
	st.stack = st.stack[:len(st.stack) - 1]
}