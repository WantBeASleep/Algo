package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type leaf struct {
	name string
	parent *leaf
}

func main() {
	cin := bufio.NewReader(os.Stdin)
	sc := bufio.NewScanner(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var n int
	fmt.Fscan(cin, &n)

	m := make(map[string]*leaf)
	for i := 0; i != n - 1; i++ {
		var parr, child string
		fmt.Fscan(cin, &child, &parr)

		_, flagParr := m[parr]
		_, flagChild := m[child]

		if flagParr {
			if flagChild {
				m[child].parent = m[parr]
			} else {
				newchild := leaf{name: child, parent: m[parr]}
				
				m[child] = &newchild
			}
		} else {
			newparr := leaf{name: parr}
			m[parr] = &newparr

			if flagChild {
				m[child].parent = &newparr
			} else {
				newchild := leaf{name: child, parent: &newparr}
				
				m[child] = &newchild
			}
		}
	}

	for sc.Scan() {
		if sc.Text() == "" {
			break
		}

		str := strings.Split(sc.Text(), " ")
		name1 := str[0]
		name2 := str[1]

		var it *leaf

		it = m[name1]
		for it != nil && it.name != name2 {
			it = it.parent
		}
		if it != nil {
			fmt.Fprintf(cout, "%d ", 2)
			continue
		}

		it = m[name2]
		for it != nil && it.name != name1 {
			it = it.parent
		}
		if it != nil {
			fmt.Fprintf(cout, "%d ", 1)
			continue
		}

		fmt.Fprintf(cout, "%d ", 0)
	}
	
	fmt.Fprintln(cout)
	cout.Flush()
}