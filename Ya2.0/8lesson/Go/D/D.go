package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var N int
	fmt.Fscan(cin, &N)

	tree := make(map[int][]int, N)
	for i := 0; i != N-1; i++ {
		var num1, num2 int
		fmt.Fscan(cin, &num1, &num2)

		tree[num1] = append(tree[num1], num2)
		tree[num2] = append(tree[num2], num1)
	}

	var ans int
	numLen(tree, 1, 1, &ans)

	fmt.Fprintln(cout, ans)
	
	cout.Flush()
}

type pathOfVertex struct{
	min, max int
}

func (p *pathOfVertex) add(x int) {
	if x > p.max {
		p.min = p.max
		p.max = x
	} else if x > p.min {
		p.min = x
	}
}

func (p *pathOfVertex) getLen() int {
	return p.max + p.min + 1
}

func numLen(tree map[int][]int, x, ancs int, maxPathLen *int) int {
	if len(tree[x]) == 0 {
		return 1
	}

	var vertx pathOfVertex

	for i := 0; i != len(tree[x]); i++ {
		if tree[x][i] == ancs {
			continue
		}
		vertx.add(numLen(tree, tree[x][i], x, maxPathLen))
		if vertx.getLen() > *maxPathLen {
			*maxPathLen = vertx.getLen()
		}
	}
	return vertx.max + 1
}