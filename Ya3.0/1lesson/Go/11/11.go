package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cin = bufio.NewScanner(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	stack := make([]int, 0)

	for cin.Scan() && cin.Text() != "exit" {
		command := strings.Split(cin.Text(), " ")
		switch command[0] {
		case "push":
			num, _ := strconv.Atoi(command[1])
			stack = push(stack, num)
			fmt.Fprintln(cout, "ok")

		case "pop":
			var num int
			var e error
			stack, num, e = pop(stack)
			if e != nil {
				fmt.Fprintln(cout, e)
			} else {
				fmt.Fprintln(cout, num)
			}
		case "back":
			num, e := back(stack)
			if e != nil {
				fmt.Fprintln(cout, e)
			} else {
				fmt.Fprintln(cout, num)
			}
		case "size":
			fmt.Fprintln(cout, size(stack))
		case "clear":
			stack = clear(stack)
			fmt.Fprintln(cout, "ok")
		}
	}

	fmt.Fprintln(cout, "bye")
	cout.Flush()
}

func push(stack []int, n int) []int {
	stack = append(stack, n)
	return stack
}

func pop(stack []int) ([]int, int, error) {
	if len(stack) == 0 {
		return stack, 0, errors.New("error")
	}

	num := stack[len(stack) - 1]
	stack = stack[:len(stack) - 1]

	return stack, num, nil
}

func back(stack []int) (int, error) {
	if len(stack) == 0 {
		return 0, errors.New("error")
	}

	return stack[len(stack) - 1], nil
}

func size(stack []int) int {
	return len(stack)
}

func clear(stack []int) []int {
	stack = stack[:0]
	return stack
}
