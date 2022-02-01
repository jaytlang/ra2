package main

import (
	"fmt"
	"testing"
)

func TestSampleGraph(t *testing.T) {
	for i := 0; i < 100; i++ {
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
		_, ainc := g.findPF(s, a, flow)
		_, ainc2 := g.findPF(d, a, flow)
		_, aout := g.findPF(a, b, flow)
		_, aout2 := g.findPF(a, c, flow)
		if ainc+ainc2 != aout+aout2 {
			t.Errorf("%d + %d != %d + %d", ainc, ainc2, aout, aout2)
		} else {
			fmt.Printf("%d + %d = %d + %d\n", ainc, ainc2, aout, aout2)
		}
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
	}
}
