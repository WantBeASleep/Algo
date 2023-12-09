package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	fmt.Fprintln(cout, foo(func(a, b int) bool {return a == b}))
	fmt.Fprintln(cout, foo(bar))
	
	cout.Flush()
}

func foo(check func(int, int) bool) bool {
	a := 5
	b := 6
	return check(a, b)
}

func bar(a, b int) bool {
	return a < b
}