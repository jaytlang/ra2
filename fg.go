package main

/* MARK: Queueing for BFS */

type q struct {
	nodes []node
	flows []int
}

func newQ() *q {
	nq := q{
		nodes: make([]node, 0),
		flows: make([]int, 0),
	}
	return &nq
}

func (q *q) push(n node, f int) {
	q.nodes = append(q.nodes, n)
	q.flows = append(q.flows, f)
}

func (q *q) pop() (node, int) {
	n := q.nodes[0]
	f := q.flows[0]
	q.nodes = q.nodes[1:]
	q.flows = q.flows[1:]
	return n, f
}

func (q *q) empty() bool {
	return len(q.nodes) == 0
}

/* MARK: Flow Graph */

type fg struct {
	ogCapacity map[node]map[node]int
	capacity   map[node]map[node]int
	adjacency  map[node][]node
}

func newFg() *fg {
	g := fg{
		ogCapacity: make(map[node]map[node]int),
		capacity:   make(map[node]map[node]int),
		adjacency:  make(map[node][]node),
	}
	return &g
}

func (g *fg) addNode(n node) {
	g.ogCapacity[n] = make(map[node]int)
	g.capacity[n] = make(map[node]int)
	g.adjacency[n] = make([]node, 0)
}

func (g *fg) addEdge(from, to node, cap int) {
	if _, ok := g.capacity[from]; !ok {
		g.addNode(from)
	}
	if _, ok := g.capacity[to]; !ok {
		g.addNode(to)
	}
	g.ogCapacity[from][to] = cap
	g.capacity[from][to] = cap
	g.adjacency[from] = append(g.adjacency[from], to)
}

type pt int

const (
	resid pt = iota
	flow
)

// node path is returned in reverse order
func (g *fg) findPF(s, e node, t pt) ([]node, int) {
	q := newQ()
	q.push(s, int(^uint(0)>>1))

	ancestry := make(map[node]node)

	for !q.empty() {
		n, f := q.pop()

		if n == e {
			e2s := make([]node, 0)
			for n := e; n != s; n = ancestry[n] {
				e2s = append(e2s, n)
			}

			return append(e2s, s), f
		}

		for _, adj := range g.adjacency[n] {
			var cf int

			if t == flow {
				cf = g.ogCapacity[n][adj] - g.capacity[n][adj]
			} else {
				cf = g.capacity[n][adj]
			}

			if cf > 0 {
				if _, ok := ancestry[adj]; !ok {
					ancestry[adj] = n
					if f < cf {
						q.push(adj, f)
					} else {
						q.push(adj, cf)
					}
				}
			}
		}
	}

	return make([]node, 0), 0
}

func (g *fg) maximize(src, sink node) int {
	rvp, f := g.findPF(src, sink, resid)
	total := 0

	for ; f > 0; rvp, f = g.findPF(src, sink, resid) {
		total += f
		n := sink
		for n != src {
			pn := rvp[0]
			rvp = rvp[1:]

			g.capacity[pn][n] -= f
			g.capacity[n][pn] += f
			n = pn
		}
	}

	return total
}
