package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// line, _ := bufio.NewReaderSize(os.Stdin, 100000).ReadString('\n')
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	list := strings.Split(line[:len(line) - 1], " ")

	m := make(map[int]int)

	for _, x := range list {
		num, _ := strconv.Atoi(x)
		m[num]++
	}

	for _, x := range list {
		num, _ := strconv.Atoi(x)
		if m[num] == 1 {
			fmt.Print(num, " ")
		}
	}

	fmt.Println()

}