// ловит TL, асимптотика такая же, только больше на 1 проход, в теории не должно TLиться
// если на плюсах проще ебнуть map'у, и добавлять за logN, то тут мы после берем ключи мапы, это o(n) и сортим их n*log(n)
// проблем быть не должно, но TL ебанутый

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[int64]int64)
	var l []int64

	for i := 0; i != n; i++ {
		var col, num int64
		fmt.Scan(&col, &num)

		_, flag := m[col]
		if !flag {
			l = append(l, col)
		}
		m[col] += num
	}

	sort.Slice(l, func (i, j int) bool {return l[i] < l[j]})

	for i := 0; i < len(l); i++ {
		fmt.Println(l[i], m[l[i]])
	}
}