package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	root *node
	cnt int
}

type node struct {
	val int
	rn *node
	ln *node
}

func (tr *tree) find(x int, cout *bufio.Writer) {
	if find(tr.root, x) {
		fmt.Fprintln(cout, "YES")
	} else {
		fmt.Fprintln(cout, "NO")
	}
}

func (tr *tree) add(x int, cout *bufio.Writer) {
	if find(tr.root, x) {
		fmt.Fprintln(cout, "ALREADY")
	} else {
		tr.root = add(tr.root, x)
		tr.cnt++
		fmt.Fprintln(cout, "DONE")
	}
}

func (tr *tree) printtree(cout *bufio.Writer) {
	printtree(tr.root, 0, cout)
}

func find(nd *node, x int) bool {
	if nd == nil {
		return false
	}
	if x == nd.val {
		return true
	}
	if x > nd.val {
		return find(nd.rn, x)
	}
	if x < nd.val {
		return find(nd.ln, x)
	}

	return false //default
}

func add(nd *node, x int) *node {
	if nd == nil {
		newnode := &node{val: x, rn: nil, ln: nil}
		return newnode
	}
	if x > nd.val {
		nd.rn = add(nd.rn, x)
	}
	if x < nd.val {
		nd.ln = add(nd.ln, x)
	}
	return nd
}

func printtree(nd *node, deep int, cout *bufio.Writer) {
	if nd == nil {
		return
	}

	printtree(nd.ln, deep + 1, cout)
	fmt.Fprintf(cout, "%s%d\n", strings.Repeat(".", deep), nd.val)
	printtree(nd.rn, deep + 1, cout)
}

func main() {
	// cin := bufio.NewReader(os.Stdin)
	cin := bufio.NewScanner(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)

	binTree := tree{root: nil, cnt: 0}
	
	for cin.Scan() {
		if cin.Text() == "" {
			break
		}

		str := strings.Split(cin.Text(), " ")

		
		switch str[0] {
		case "ADD":
			num, _ := strconv.Atoi(str[1])
			binTree.add(num, cout)
		case "SEARCH":
			num, _ := strconv.Atoi(str[1])
			binTree.find(num, cout)
		case "PRINTTREE":
			binTree.printtree(cout)
		}

	}



	
	cout.Flush()
}