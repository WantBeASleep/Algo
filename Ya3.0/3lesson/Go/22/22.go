package main

import (
	"fmt"
	"bufio"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

func main() {
	var n, k int
	fmt.Fscan(cin, &n, &k)

	
	cout.Flush()
}