package main

import (
	"bufio"
	"sort"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)
	
	for reader.Scan() {
		if (reader.Text() == "") {
			break
		}
		list := strings.Split(reader.Text(), " ")
		num, _ := strconv.Atoi(list[1])
		m[list[0]] += num
	}

	list := make([]string, len(m))
	var i int
	for k := range m {
		list[i] = k
		i++
	}
	sort.Strings(list)

	for _, s := range list {
		fmt.Println(s, m[s])
	}

}