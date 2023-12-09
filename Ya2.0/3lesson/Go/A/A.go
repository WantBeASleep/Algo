package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100000)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')

	strlist1 := strings.Split(line1[:len(line1) - 1], " ")
	strlist2 := strings.Split(line2[:len(line2) - 1], " ")

	m := make(map[int]int)
	

	for _, s := range strlist1 {
		num, _ := strconv.Atoi(s)
		m[num] = 1
	}

	var count int

	for _, s := range strlist2 {
		num, _ := strconv.Atoi(s)
		_, flag := m[num]
		if flag {
			count++
		}
	}

	fmt.Println(count)
}