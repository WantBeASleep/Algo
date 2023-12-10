package main

import (
	"bufio"
	"fmt"
	"math"
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
		var tin, time int
		fmt.Fscan(cin, &tin, &time)

		// 1 поступил груз
		// -1 обработан груз
		events[2 * i] = event{tin, 1}
		events[2 * i + 1] = event{tin + time, -1}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].status < events[j].status
		} else {
			return events[i].time < events[j].time
		}
	})

	// fmt.Fprintln(cout, events)

	var maxload int
	var curload int
	for _, ev := range events {
		if ev.status == 1 {
			curload++
			maxload = int(math.Max(float64(maxload), float64(curload)))
		} else {
			curload--
		}
	}

	fmt.Fprintln(cout, maxload)
	
	cout.Flush()
}