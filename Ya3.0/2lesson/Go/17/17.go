package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	q1 := queue{data: make([]int, 5)}
	q2 := queue{data: make([]int, 5)}

	fmt.Fscanf(cin, "%d %d %d %d %d\n", &q1.data[0], &q1.data[1], &q1.data[2], &q1.data[3], &q1.data[4])
	fmt.Fscanf(cin, "%d %d %d %d %d\n", &q2.data[0], &q2.data[1], &q2.data[2], &q2.data[3], &q2.data[4])
	// fmt.Fprintf(cout, "%v:%v\n", q1.data, q2.data)

	for i := 0; i != int(1e6); i++ {

		// fmt.Fprintf(cout,
		// "Step:%d\nf: %v\ns: %v\n",
		// i + 1, q1.data, q2.data)


		first, second := q1.pop(), q2.pop()

		if (first > second && !(first == 9 && second == 0)) || (first == 0 && second == 9) {
			// fmt.Fprintf(cout, "fWin step:%d, f:%d, s:%d\n", i + 1, first, second)
			q1.push(first)
			q1.push(second)
			
			if q2.size() == 0 {
				fmt.Fprintln(cout, "first", i + 1)
				cout.Flush()
				os.Exit(0)
			}
		} else {
			// fmt.Fprintf(cout, "sWin step:%d, f:%d, s:%d\n", i + 1, first, second)
			q2.push(first)
			q2.push(second)

			if q1.size() == 0 {
				fmt.Fprintln(cout, "second", i + 1)
				cout.Flush()
				os.Exit(0)
			}
		}
	}

	fmt.Fprintln(cout, "botva")
	cout.Flush()
}

type queue struct {
	data []int
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) pop() int {
	ans := q.data[0]
	q.data = q.data[1:]
	return ans
}

func (q *queue) push(x int) {
	q.data = append(q.data, x)
}