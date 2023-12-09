package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)

	var str string
	fmt.Fscan(cin, &str)

	preff := make([]int, len(str) + 1)

	flagOfWrong := false

	for i := 1; i != len(preff); i++ {
		if str[i - 1] == '(' {
			preff[i] = preff[i - 1] + 1
		} else {
			preff[i] = preff[i - 1] - 1
		}
		
		if preff[i] < 0 {
			flagOfWrong = true
		}
	}

	if flagOfWrong || preff[len(preff) - 1] != 0 {
		fmt.Fprintln(cout, "NO")
	} else {
		fmt.Fprintln(cout, "YES")
	}

	cout.Flush()
}