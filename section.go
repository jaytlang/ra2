package main

import "fmt"

type section struct {
	isTutorial bool
	time       st
	instructor string
}

func (sec section) String() string {
	if sec.isTutorial {
		return fmt.Sprintf("Tutorial %s: taught by %s\n", sec.time, sec.instructor)
	} else {
		return fmt.Sprintf("Recitation %s: taught by %s\n", sec.time, sec.instructor)
	}
}

/* MARK: section-related utilities */

func findSectNs(inst string, rec bool) []int {
	m := tbs
	if rec {
		m = rbs
	}

	sns := []int{}
	for sn, s := range m {
		if s.instructor == inst {
			sns = append(sns, sn)
		}
	}

	return sns
}

func legalRTSectPairs() [][2]int {
	pairs := make([][2]int, 0)

	for ri, tis := range r2t {
		for _, ti := range tis {
			rsns := findSectNs(ri, true)
			tsns := findSectNs(ti, false)
			for _, rsn := range rsns {
				for _, tsn := range tsns {
					pairs = append(pairs, [2]int{rsn, tsn})
				}
			}
		}
	}

	return pairs
}
