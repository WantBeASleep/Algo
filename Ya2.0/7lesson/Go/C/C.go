package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type section struct {
	start int
	end int
}

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var m int
	fmt.Fscan(cin, &m)

	
	var l, r int
	var sections []section
	fmt.Fscan(cin, &l, &r)
	for !(l == 0 && r == 0) {
		if !(r <= 0 || l >= m) {
			sections = append(sections, section{l, r})
		}
	}

	sort.Slice(sections, func(i, j int) bool {
		if sections[i].start == sections[j].start {
			return sections[i].end < sections[j].end
		} else {
			return sections[i].start < sections[j].start
		}
	})

	curend := sections[0].end
	var curl, curr int
	var ans []section
	for _, sect := range sections {
		if sect.start > curend {
			curend = curr
			ans = append(ans, section{curl, curr})
			if curend >= m {
				break
			}
		}

		if sect.end <= curr {
			continue
		}
		if sect.start <= curend && sect.end > curr{
			curl = sect.start
			curr = sect.end
		}

	}
	
	cout.Flush()
}