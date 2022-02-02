package main

import "github.com/yourbasic/graph"

type fg struct {
	tl map[*fn]int
	g  *graph.Mutable
}

type fgr struct {
	tl   map[*fn]int
	flow int
	fg   graph.Iterator
}

/* MARK: flow graphs */

func newFg(fns []*fn) *fg {
	f := fg{
		tl: make(map[*fn]int),
		g:  graph.New(len(fns)),
	}

	for i, fn := range fns {
		f.tl[fn] = i
	}

	return &f
}

func (f *fg) addCost(fn1, fn2 *fn, cost int) {
	f.g.AddCost(f.tl[fn1], f.tl[fn2], int64(cost))
}

func (f *fg) runMaxFlow(s, e *fn) *fgr {
	flow, nfg := graph.MaxFlow(f.g, f.tl[s], f.tl[e])
	r := fgr{
		tl:   f.tl,
		flow: int(flow),
		fg:   nfg,
	}

	return &r
}

/* MARK: flow graph results */
func (f *fgr) flowAlong(s, e *fn) int {
	target := f.tl[e]
	res := 0
	f.fg.Visit(f.tl[s], func(t int, c int64) bool {
		if t == target {
			res = int(c)
			return true
		}
		return false
	})

	return res
}
