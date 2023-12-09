package main
/* 
можно написать сильно лучше, но пока не понял как
я ввожу свидителей просто в срез
затем ввожу номер, кидаю его в срез, делаю из него посимвольный map
потом итерирую свидителей, итерирую посимвольно каждого свидителя и сравниваю его с текущим номером
запоминаю его совпадение в мапе и потом 2-ым проходом вывожу все номера

он делает тоже самое, только делая set еще и по свидителям и сравнивая на вхождения сеты
уверен в го так как нибудь можно, но я не особо искал
*/
import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)

	witness := make([]string, m)
	for i := 0; i != m; i++ {
		fmt.Scan(&witness[i])
	}

	var n int
	fmt.Scan(&n)

	var coincidence int
	order := make([]string, n)
	unorder := make(map[string]int)
	for i := 0; i != n; i++ {
		fmt.Scan(&order[i])

		car := make(map[rune]int)
		for _, char := range order[i] {
			car[char] = 1
		}

		var curcoincidence int

		for j := 0; j != m; j++ {
			flag := true
			for _, char := range witness[j] {
				_, res := car[char]
				if !res {
					flag = false
				}
			}

			if flag {
				curcoincidence++
			}
		}

		if curcoincidence > coincidence {
			coincidence = curcoincidence
		}

		unorder[order[i]] = curcoincidence
	}

	for i := 0; i != n; i++ {
		if unorder[order[i]] == coincidence {
			fmt.Println(order[i])
		}
	}
}