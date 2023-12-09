package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type wrdscnt struct {
	cnt int
	wrd string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)

	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		line := strings.Split(sc.Text(), " ")
		for _, s := range line {
			m[s]++
		}
	}

	ans := make([]wrdscnt, len(m))
	var i int
	for k, v := range m {
		ans[i] = wrdscnt{v, k}
		i++
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i].cnt == ans[j].cnt {
			return ans[i].wrd < ans[j].wrd
		} else {
			return ans[i].cnt > ans[j].cnt
		}
	})

	for _, s := range ans {
		fmt.Println(s.wrd)
	}
}