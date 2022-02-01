package main

import (
	"fmt"
)

const (
	kid = iota
	rtunion
	tutorial
	src
	sink
)

type fn struct {
	t int

	st   *student
	rsec *section
	tsec *section
}

func (f fn) String() string {
	switch f.t {
	case kid:
		return f.st.String()
	case rtunion:
		return fmt.Sprintf("Union of:\n\t%s\n\t%s\n", f.rsec.String(), f.tsec.String())
	case tutorial:
		return f.tsec.String()
	case src:
		return "Source"
	default:
		return "Sink"
	}
}

func makeStudentFns() ([]*fn, error) {
	s, err := newStudentsAll()
	if err != nil {
		return nil, err
	}

	fns := make([]*fn, len(s))
	for i, ts := range s {
		fns[i] = &fn{
			t:  kid,
			st: ts,
		}
	}
	return fns, nil
}

func makeTutFns() ([]*fn, error) {
	fns := make([]*fn, 0)
	for _, tb := range tbs {
		fns = append(fns, &fn{
			t:    tutorial,
			tsec: tb,
		})
	}

	return fns, nil
}

func makeRTUFns() ([]*fn, error) {
	fns := make([]*fn, 0)
	lrts := legalRTSectPairs()

	for _, rt := range lrts {
		fns = append(fns, &fn{
			t:    rtunion,
			rsec: rbs[rt[0]],
			tsec: tbs[rt[1]],
		})
	}

	return fns, nil
}
