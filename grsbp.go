package main

// There is an O(N^3) algorithm for this problem (3D-SRP-SYM-BIN),
// I'll get there at some point. We need to make sure
// this works at all first.

type sbp struct {
	adjacency map[*fn][]*fn
	perfects  int
	pairs     int
	loners    int
}

func newSbp() *sbp {
	return &sbp{
		adjacency: make(map[*fn][]*fn),
	}
}

func (s *sbp) addNode(n *fn) {
	s.adjacency[n] = make([]*fn, 0)
}

func (s *sbp) addPref(from, to *fn) {
	if _, ok := s.adjacency[from]; !ok {
		s.addNode(from)
	}
	if _, ok := s.adjacency[to]; !ok {
		s.addNode(to)
	}
	s.adjacency[from] = append(s.adjacency[from], to)
}

func (s *sbp) claim(n *fn) {
	for _, n2 := range s.adjacency[n] {
		for i, el := range s.adjacency[n2] {
			if el == n {
				s.adjacency[n2] = append(s.adjacency[n2][:i], s.adjacency[n2][i+1:]...)
				break
			}
		}
	}

	delete(s.adjacency, n)
}

func (s *sbp) compatible(a, b *fn) bool {
	w1 := false
	w2 := false

	for _, n := range s.adjacency[a] {
		if n == b {
			w1 = true
			break
		}
	}
	for _, n := range s.adjacency[b] {
		if n == a {
			w2 = true
			break
		}
	}

	return w1 && w2
}

func (s *sbp) matchThrees() map[*fn][]*fn {
	m := make(map[*fn][]*fn)
	a := s.adjacency

	claimed := make(map[*fn]bool)

	for i := range a {
		for j := range a {
			for k := range a {
				if i == j || i == k || j == k {
					continue
				}
				if s.compatible(i, j) && s.compatible(i, k) && s.compatible(j, k) {
					if !claimed[i] && !claimed[j] && !claimed[k] {
						m[i] = []*fn{j, k}
						m[j] = []*fn{i, k}
						m[k] = []*fn{i, j}
						claimed[i] = true
						claimed[j] = true
						claimed[k] = true
						s.perfects += 3
					}
				}
			}
		}
	}

	for k := range claimed {
		s.claim(k)
	}

	return m
}

func (s *sbp) getAllPairs() [][]*fn {
	pairs := make([][]*fn, 0)
	a := s.adjacency

	for i := range a {
		for j := range a {
			if i == j {
				continue
			} else if s.compatible(i, j) {
				pairs = append(pairs, []*fn{i, j})
			}
		}
	}

	return pairs
}

func (s *sbp) getAllRemaining() []*fn {
	l := make([]*fn, 0)

	for i := range s.adjacency {
		l = append(l, i)
	}
	return l
}

func (s *sbp) matchPairsAndLoners() map[*fn][]*fn {
	m := make(map[*fn][]*fn)

	for {
		p := s.getAllPairs()
		if len(p) == 0 {
			return m
		}

		cp := p[0]
		s.claim(cp[0])
		s.claim(cp[1])

		l := s.getAllRemaining()

		if len(l) == 0 {
			s.addPref(cp[0], cp[1])
			s.addPref(cp[1], cp[0])
			return m
		}

		cl := l[0]
		s.claim(cl)

		m[cp[0]] = []*fn{cp[1], cl}
		m[cp[1]] = []*fn{cp[0], cl}
		m[cl] = []*fn{cp[0], cp[1]}
		s.pairs += 2
		s.loners += 1
	}
}

func (s *sbp) matchLonerOnlyTeams() map[*fn][]*fn {
	m := make(map[*fn][]*fn)

	for {
		l := s.getAllRemaining()
		if len(l) < 3 {
			break
		}

		cl := l[0:3]
		s.claim(cl[0])
		s.claim(cl[1])
		s.claim(cl[2])
		s.loners += 3

		m[cl[0]] = []*fn{cl[1], cl[2]}
		m[cl[1]] = []*fn{cl[0], cl[2]}
		m[cl[2]] = []*fn{cl[0], cl[1]}
	}

	return m
}

func (s *sbp) augmentMatching(m *map[*fn][]*fn, maxSz int) {
	l := s.getAllRemaining()
	for _, n := range l {
		for t1, tr := range *m {
			if len(tr) <= (maxSz - 1) {
				s.claim(n)
				s.loners += 1

				(*m)[n] = append(tr, t1)
				(*m)[t1] = append(tr, n)
				for _, t2 := range tr {
					(*m)[t2] = append((*m)[t2], n)
				}
				break
			}
		}
	}

}

func (s *sbp) greedyMatching() map[*fn][]*fn {
	m := s.matchThrees()
	for k, v := range s.matchPairsAndLoners() {
		m[k] = v
	}

	if len(s.getAllRemaining()) >= 3 {
		for k, v := range s.matchLonerOnlyTeams() {
			m[k] = v
		}
	}

	s.augmentMatching(&m, 3)
	s.augmentMatching(&m, 4)

	return m
}
