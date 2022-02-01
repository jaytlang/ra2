package main

import "testing"

func TestSampleGraph(t *testing.T) {
	g := newFg()

	var (
		s node = new(interface{})
		a node = new(interface{})
		b node = new(interface{})
		c node = new(interface{})
		d node = new(interface{})
		e node = new(interface{})
	)

	g.addEdge(s, a, 7)
	g.addEdge(s, d, 4)
	g.addEdge(a, b, 5)
	g.addEdge(a, c, 3)
	g.addEdge(b, e, 8)
	g.addEdge(c, b, 3)
	g.addEdge(c, e, 5)
	g.addEdge(d, a, 3)
	g.addEdge(d, c, 2)

	if g.maximize(s, e) != 10 {
		t.Error("Expected 10, got", g.maximize(s, e))
	}
}

func TestAnotherSampleGraph(t *testing.T) {
	g := newFg()

	var (
		s node = new(interface{})
		a node = new(interface{})
		b node = new(interface{})
		e node = new(interface{})
	)

	g.addEdge(s, a, 5)
	g.addEdge(s, b, 5)
	g.addEdge(a, e, 5)
	g.addEdge(b, e, 5)
	g.addEdge(a, b, 10)

	if g.maximize(s, e) != 10 {
		t.Error("Expected 10, got", g.maximize(s, e))
	} else if _, f := g.findPF(s, e, flow); f != 5 {
		t.Error("Expected 5, got", f)
	} else if _, f := g.findPF(a, b, flow); f != 0 {
		t.Error("Expected 0, got", f)
	} else if _, f := g.findPF(s, a, flow); f != 5 {
		t.Error("Expected 5, got", f)
	}
}
