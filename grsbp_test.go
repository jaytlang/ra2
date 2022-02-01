package main

import (
	"fmt"
	"testing"
)

var (
	a node = new(interface{})
	b node = new(interface{})
	c node = new(interface{})
	d node = new(interface{})
	e node = new(interface{})
	f node = new(interface{})
	g node = new(interface{})
)

// These just ensure everything looks ok via console

func TestEvenMatching(t *testing.T) {
	s := newSbp()

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(b, c)
	s.addPref(c, b)
	s.addPref(a, c)
	s.addPref(c, a)

	s.addPref(d, e)
	s.addPref(e, d)
	s.addPref(e, f)
	s.addPref(f, e)
	s.addPref(d, f)
	s.addPref(f, d)

	m := s.greedyMatching()
	fmt.Println(m)
}

func TestPairwiseMatchingNoLoners(t *testing.T) {
	s := newSbp()

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(d, e)
	s.addPref(e, d)

	m := s.greedyMatching()
	fmt.Println(m)
}

func TestPairwiseMatchingLoner(t *testing.T) {
	s := newSbp()

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(b, c)
	s.addPref(c, b)
	s.addPref(a, c)
	s.addPref(c, a)

	s.addPref(d, e)
	s.addPref(e, d)

	s.addNode(f)
	m := s.greedyMatching()
	fmt.Println(m)
}

func TestManyLoners(t *testing.T) {
	s := newSbp()

	s.addNode(a)
	s.addNode(b)
	s.addNode(c)
	s.addNode(d)
	s.addNode(e)
	s.addNode(f)

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(b, c)
	s.addPref(c, b)
	s.addPref(a, c)
	s.addPref(c, a)

	m := s.greedyMatching()
	fmt.Println(m)
}

func TestRunOutOfLonersFive(t *testing.T) {
	s := newSbp()

	s.addNode(a)
	s.addNode(b)
	s.addNode(c)
	s.addNode(d)
	s.addNode(e)

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(c, d)
	s.addPref(d, c)

	m := s.greedyMatching()
	fmt.Println(m)

}

func TestRunOutOfLoners(t *testing.T) {
	s := newSbp()

	s.addNode(a)
	s.addNode(b)
	s.addNode(c)
	s.addNode(d)
	s.addNode(e)
	s.addNode(f)
	s.addNode(g)

	s.addPref(a, b)
	s.addPref(b, a)
	s.addPref(c, d)
	s.addPref(d, c)
	s.addPref(e, f)
	s.addPref(f, e)

	m := s.greedyMatching()
	fmt.Println(m)

}
