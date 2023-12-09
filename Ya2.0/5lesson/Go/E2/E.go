package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	cout := bufio.NewWriter(os.Stdout)
	
	var S int
	fmt.Fscan(cin, &S)

	var sz1, sz2, sz3 int

	fmt.Fscan(cin, &sz1)
	A := make([]int, sz1)
	for i := 0; i != sz1; i++ {
		fmt.Fscan(cin, &A[i])
	}

	fmt.Fscan(cin, &sz2)
	B := make([]int, sz2)
	for i := 0; i != sz2; i++ {
		fmt.Fscan(cin, &B[i])
	}

	fmt.Fscan(cin, &sz3)
	C := make(map[int]int, sz3)
	for i := 0; i != sz3; i++ {
		var num int
		fmt.Fscan(cin, &num)

		_, flag := C[num]
		if !flag {
			C[num] = i
		}
	}

	ansflag := false
	for i := 0; i != sz1; i++ {
		if ansflag {
			break
		}
		for j := 0; j != sz2; j++ {
			if ansflag {
				break
			}
			mod := S - A[i] - B[j]
			_, flag := C[mod]

			if flag {
				fmt.Fprintln(cout, i, j, C[mod])
				ansflag = true
			}
		}
	}

	if !ansflag {
		fmt.Fprintln(cout, -1)
	}
	
	cout.Flush()
}