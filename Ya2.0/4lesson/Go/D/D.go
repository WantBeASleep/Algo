package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

type mod struct {
	frac float64
	name string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var total int
	condition := make(map[string]int)
	var order []string
	for (sc.Scan()) {
		if sc.Text() == "" {
			break
		}

		pos := strings.LastIndex(sc.Text(), " ")
		partyname := sc.Text()[:pos]
		votes, _ := strconv.Atoi(sc.Text()[pos + 1:])

		_, flag := condition[partyname]
		if !flag {
			order = append(order, partyname)
		}
		condition[partyname] += votes
		total += votes
	}

	firstConst := float64(total) / 450

	ans := make(map[string]int)
	modlist := make([]mod, len(order))

	var firstDistrub int
	var i int
	for _, s := range order {
		stats := float64(condition[s]) / firstConst
		whol, frac := math.Modf(stats)

		ans[s] = int(whol)
		firstDistrub += ans[s]

		modlist[i] = mod{frac, s}
		i++
	}

	if (firstDistrub != 450) {
		sort.Slice(modlist, func(i, j int) bool {
			if math.Abs(modlist[i].frac - modlist[j].frac) < 0.00001 {
				return condition[modlist[i].name] > condition[modlist[j].name]
			} else {
				return modlist[i].frac > modlist[j].frac
			}
		})

		var i int
		for firstDistrub != 450 {
			i %= len(modlist)
			ans[modlist[i].name]++
			firstDistrub++
			i++
		}
	}

	for _, s := range order {
		fmt.Println(s, ans[s])
	}
}