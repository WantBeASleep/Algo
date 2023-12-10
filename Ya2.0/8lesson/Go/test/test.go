package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// cin := bufio.NewReader(os.Stdin)
	cin := bufio.NewScanner(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	cin.Scan()
	fmt.Fprint(cout, cin.Text())
	
	
	cout.Flush()
}