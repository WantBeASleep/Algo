package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var cin = bufio.NewScanner(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {

	q := queue{data: make([]int, 0, 8)}
	for cin.Scan() && cin.Text() != "exit" {
		switch cin.Text() {
		case "pop":
			ans, e := q.pop()
			if e != nil {
				fmt.Fprintln(cout, "error")
			} else {
				fmt.Fprintln(cout, ans)
			}
		case "front":
			ans, e := q.front()
			if e != nil {
				fmt.Fprintln(cout, "error")
			} else {
				fmt.Fprintln(cout, ans)
			}
		case "size":
			fmt.Fprintln(cout, q.size())
		case "clear":
			q.clear()
			fmt.Fprintln(cout, "ok")
		default:
			args := strings.Split(cin.Text(), " ")
			num, _ := strconv.Atoi(args[1])
			q.push(num)
			fmt.Fprintln(cout, "ok")
		}
	}

	fmt.Fprintln(cout, "bye")
	cout.Flush()
}

type queue struct {
	data []int
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) clear() {
	q.data = make([]int, 0, 8)
}

func (q *queue) front() (int, error) {
	if len(q.data) == 0 {
		return 0, errors.New("queue empty")
	}

	return q.data[0], nil
}


func (q *queue) pop() (int, error) {
	ans, e := q.front()
	if e != nil {
		return 0, errors.New("queue empty")
	} else {
		q.data = q.data[1:]
	}
	return ans, nil
}

func (q *queue) push(x int) {
	q.data = append(q.data, x)
}