/*
задача чисто помериться членом
пиздец громоздкая и неподъемная

если вкратце, то минимум из максимума подбирается бинаркой с ограном сверху, вопросов 0

Далее нужно проверить что добираемся в n за d щагов с ограном бинарки
В графе нет циклов, пихать сюда дейкстру или флойда особо смысла нет, просто обход в ширину

обход в ширину ~O(N + M) + бинарка, итого O((N + M)Log(N?))

Ошибка в 44 тесте, честно, ни сил, ни малейшего желания сейчас её искать, я просто пиздец как so tired

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

type edge struct {
	from, to, num int
}

func main() {
	var n, m, d int
	fmt.Fscan(cin, &n, &m, &d)

	graph := make(map[int][]edge, n)
	for i := 0; i != m; i++ {
		newEdge := edge{}
		fmt.Fscan(cin, &newEdge.from, &newEdge.to, &newEdge.num)
		graph[newEdge.from] = append(graph[newEdge.from], newEdge)
	}

	maxN := bin(graph, n, d)

	if maxN == -1 {
		fmt.Fprintln(cout, -1)
	} else {
		path := restorePath(graph, n, maxN)
	
		fmt.Fprintln(cout, len(path) - 1)
		for i := len(path) - 1; i != -1; i-- {
			fmt.Fprintf(cout, "%d ", path[i])
		}
		fmt.Fprintln(cout)
	}
	
	cout.Flush()
}

func restorePath(graph map[int][]edge, n, weightLimit int) []int {
	techPath := getPath(graph, n, weightLimit)

	var path []int
	
	for i := n; i != -1; i = techPath[i] {
		path = append(path, i)
	}

	return path
}

func getPath(graph map[int][]edge, n, weightLimit int) map[int]int {
	var deque []int
	path := make(map[int]int)
	visit := make(map[int]bool, len(graph))

	deque = append(deque, 1)
	path[1] = -1
	visit[1] = true

	for len(deque) > 0 && !visit[n] {
		preStepCnt := len(deque)
		for i := 0; i != preStepCnt; i++ {
			for j := 0; j != len(graph[deque[i]]); j++ {
				numVert := graph[deque[i]][j].to
				numAncs := graph[deque[i]][j].from
				numEdge := graph[deque[i]][j].num
				
				if numEdge <= weightLimit && !visit[numVert] {
					deque = append(deque, numVert)
					path[numVert] = numAncs
					visit[numVert] = true

					if numVert == n {
						return path
					}
				}
			}
		}

		deque = deque[preStepCnt:]
	}

	return path
}

func good(maxN, n, d int, graph map[int][]edge) bool {
	var deque []int
	visit := make(map[int]bool, len(graph))

	deque = append(deque, 1)
	visit[1] = true

	step := 0
	for len(deque) > 0 && !visit[n] {
		step++
		preStepCnt := len(deque)
		for i := 0; i != preStepCnt; i++ {
			for j := 0; j != len(graph[deque[i]]); j++ {
				numVert := graph[deque[i]][j].to
				numEdge := graph[deque[i]][j].num
				
				if numEdge <= maxN && !visit[numVert] {
					deque = append(deque, numVert)
					visit[numVert] = true
				}
			}
		}

		deque = deque[preStepCnt:]
	}

	return visit[n] && step <= d
}

func bin(graph map[int][]edge, n, d int) int {
	if !good(int(1e9 + 1), n, d, graph) {
		return -1
	}
	
	l := 0
	r := 1
	for !good(r, n, d, graph) && r <= int(1e9){
		r *= 2
	}


	for r != l + 1 {
		mid := (l + r) / 2
		if good(mid, n, d, graph) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}

