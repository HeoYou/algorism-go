import (
	"sort"
)

func solution(n int, edge [][]int) int {

	visit := make([]int, n+1)

	graph := make_graph(n, edge)

	var q []int
	q = append(q, 1)
	visit[1] = 1

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for i := 0; i < len(graph[node]); i++ {

			if visit[graph[node][i]] == 0 {
				visit[graph[node][i]] = visit[node] + 1
				q = append(q, graph[node][i])
			}

		}
	}
	sort.Ints(visit)
	var answer int
	for i := 0; i < n+1; i++ {
		if visit[n] == visit[i] {
			answer++
		}
	}
	return answer
}

func make_graph(n int, edge [][]int) [][]int {

	graph := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(edge); i++ {
		graph[edge[i][0]] = append(graph[edge[i][0]], edge[i][1])
		graph[edge[i][1]] = append(graph[edge[i][1]], edge[i][0])
	}

	return graph
}