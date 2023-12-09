package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := make(map[int]int)

	for i := 1; i != n + 1; i++ {
		ans[i] = 1
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		line = line[:len(line) - 1]
		
		if line == "HELP" {
			break;
		}

		status, _ := reader.ReadString('\n')
		status = status[:len(status) - 1]

		strlist := strings.Split(line, " ")
		
		if status == "YES" {
			newmap := make(map[int]int)
			for _, s := range strlist {
				num, _ := strconv.Atoi(s)
				_, flag := ans[num]
				if flag {
					newmap[num] = 1
				}
			}

			ans = newmap
		} else {
			for _, s := range strlist {
				num, _ := strconv.Atoi(s)
				_, flag := ans[num]
				if flag {
					delete(ans, num)
				}
			}
		}
	}

	var anslist []int
	for num := range ans {
		anslist = append(anslist, num)
	}

	sort.Ints(anslist)

	for _, num := range anslist {
		fmt.Print(num, " ")
	}

	fmt.Println()
}

