package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// line, _ := bufio.NewReaderSize(os.Stdin, 1000000).ReadString('\n')
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	line, _ := reader.ReadString('\n')
	list := strings.Split(line[:len(line) - 1], " ")

	m := make(map[int]int)

	for _, s := range list {
		num, _ := strconv.Atoi(s)
		_, flag := m[num]
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
			m[num] = 1
		}
	}
}