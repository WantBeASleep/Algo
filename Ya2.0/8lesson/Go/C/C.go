package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var N int
	fmt.Fscan(cin, &N)	

	tree := make(map[string]string, N)

	for i := 0; i != N - 1; i++ {
		var child, ancs string
		fmt.Fscan(cin, &child, &ancs)

		tree[child] = ancs
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		if sc.Text() == "" {
			break
		}

		str := strings.Split(sc.Text(), " ")
		name1 := str[0]
		name2 := str[1]

		path := make(map[string]int)
		it := name1
		for it != "" {
			path[it] = 1
			it = tree[it]
		}

		it = name2
		_, flagF := path[it]

		for !flagF {
			it = tree[it]
			_, flagF = path[it]
		}

		fmt.Fprintln(cout, it)
	}
	
	cout.Flush()
}