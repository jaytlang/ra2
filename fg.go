package main

/* MARK: Queueing for BFS */

type q struct {
	nodePaths [][]node
	flows     []int
}

func newQ() *q {
	nq := q{
		nodePaths: make([][]node, 0),
		flows:     make([]int, 0),
	}
	return &nq
}

func (q *q) push(ns []node, f int) {
	q.nodePaths = append(q.nodePaths, ns)
	q.flows = append(q.flows, f)
}

func (q *q) pop() ([]node, int) {
	ns := q.nodePaths[0]
	f := q.flows[0]
	q.nodePaths = q.nodePaths[1:]
	q.flows = q.flows[1:]
	return ns, f
}

func (q *q) empty() bool {
	return len(q.nodePaths) == 0
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
	q.push([]node{s}, int(^uint(0)>>1))

	for !q.empty() {
		ns, f := q.pop()
		n := ns[0]

		if n == e {
			return ns, f
		}

		for _, adj := range g.adjacency[n] {
			var cf int

			if t == flow {
				cf = g.ogCapacity[n][adj] - g.capacity[n][adj]
			} else {
				cf = g.capacity[n][adj]
			}

			nxtns := append([]node{adj}, ns...)
			if cf > 0 {
				if f < cf {
					q.push(nxtns, f)
				} else {
					q.push(nxtns, cf)
				}
			}
		}
	}

	return make([]node, 0), 0
}

func (g *fg) maximize(src, sink node) int {
	total := 0

	for {
		rvp, f := g.findPF(src, sink, resid)
		if f <= 0 {
			break
		}
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
