package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type event struct {
	time int
	status int
}

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var n int
	fmt.Fscan(cin, &n)

	events := make([]event, 2 * n)
	for i := 0; i != n; i++ {
		var tin, tout int
		fmt.Fscan(cin, &tin, &tout)

		// -1 открытие
		// 1 закрытие
		events[2 * i] = event{tin, -1}
		events[2 * i + 1] = event{tout, 1}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].status < events[j].status
		} else {
			return events[i].time < events[j].time
		}
	})

	var colorline int
	var openlines int
	for i, ev := range events {
		if openlines > 0 {
			colorline += ev.time - events[i - 1].time
		}

		if ev.status == -1 {
			openlines++
		} else {
			openlines--
		}
	}

	fmt.Fprintln(cout, colorline)
	
	cout.Flush()
}