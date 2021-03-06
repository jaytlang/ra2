package main

import "fmt"

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
	team []*fn
}

func (f fn) String() string {
	switch f.t {
	case kid:
		s := f.st.String()
		if f.rsec != nil {
			s += "Assigned to " + f.rsec.String()
			s += "Assigned to " + f.tsec.String()
		}
		if f.team != nil {
			s += "Team:"
			for _, teammate := range f.team {
				s += " " + teammate.st.email
			}
			s += "\n"
		}
		return s
	case rtunion:
		return fmt.Sprintf("Union of:\n\t%s\t%s\n", f.rsec.String(), f.tsec.String())
	case tutorial:
		return f.tsec.String()
	case src:
		return "Source\n"
	default:
		return "Sink\n"
	}
}

func makeTutFns() ([]*fn, error) {
	fns := make([]*fn, 0)
	for _, tb := range ats {
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
			rsec: ars[rt[0]],
			tsec: ats[rt[1]],
		})
	}

	return fns, nil
}
