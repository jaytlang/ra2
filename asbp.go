package main

import "fmt"

type asbp struct {
	t2sbp map[*section]*sbp
	res   map[*fn][]*fn
}

func (s *asbp) prepare(in []*fn) error {
	s.res = make(map[*fn][]*fn)
	s.t2sbp = make(map[*section]*sbp)

	s2f := make(map[*student]*fn)
	for _, sfn := range in {
		s2f[sfn.st] = sfn

		if _, ok := s.t2sbp[sfn.tsec]; !ok {
			s.t2sbp[sfn.tsec] = newSbp()
		}
		s.t2sbp[sfn.tsec].addNode(sfn)
	}

	for _, sfn := range in {
		st := sfn.st
		t := sfn.tsec
		for _, fav := range st.dpfavs {
			if s2f[fav].tsec == t {
				s.t2sbp[t].addPref(sfn, s2f[fav])
			}
		}
	}
	return nil
}

func (s *asbp) execute() error {
	for _, sbp := range s.t2sbp {
		tr := sbp.greedyMatching()
		for k, v := range tr {
			s.res[k] = v
		}
	}
	return nil
}

func (s *asbp) export() ([]*fn, error) {
	tm4 := 0
	tm5 := 0
	for _, tm := range s.res {
		if len(tm) == 3 {
			tm4++
		} else if len(tm) == 4 {
			tm5++
		}
	}

	fmt.Println("tm4: ", tm4/4)
	fmt.Println("tm5: ", tm5)
	return nil, nil
}