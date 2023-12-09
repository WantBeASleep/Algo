package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elem struct {
	count int
	idx int
}

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)

	var N, M int
	fmt.Fscan(cin, &N, &M)

	stud := make([]elem, N)
	rooms := make([]elem, M)

	for i := 0; i != N; i++ {
		var cnt int
		fmt.Fscan(cin, &cnt)

		stud[i] = elem{count: cnt, idx: i}
	}

	for i := 0; i != M; i++ {
		var cnt int
		fmt.Fscan(cin, &cnt)

		rooms[i] = elem{count: cnt, idx: i}
	}

	sort.Slice(stud, func(i, j int) bool {return stud[i].count < stud[j].count})
	sort.Slice(rooms, func(i, j int) bool {return rooms[i].count < rooms[j].count})

	ans := make([]int, N)
	var success int

	var i, j int
	for {
		if i == N || j == M {
			break
		}

		if stud[i].count < rooms[j].count {
			ans[stud[i].idx] = rooms[j].idx + 1
			i++
			j++
			success++
		} else {
			j++
		}
	}

	fmt.Fprintln(cout, success)
	for i := 0; i != N; i++ {
		fmt.Fprint(cout, ans[i], " ")
	}
	fmt.Fprintln(cout)

	cout.Flush()
}