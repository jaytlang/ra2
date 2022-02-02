package main

import "fmt"

type stats struct {
	l []*fn
}

type percent struct {
	n int
	d int
}

func (p percent) String() string {
	return fmt.Sprintf("%d%%", p.n*100/p.d)
}

func (t *stats) rptPctAssigned() {
	total := 0
	for _, fn := range t.l {
		if fn.rsec != nil && fn.tsec != nil {
			total++
		}
	}

	fmt.Printf("Percent of students assigned to a fitting section: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctBestRec() {
	total := 0
	for _, fn := range t.l {
		if fn.st.rp[0] == fn.rsec.time {
			total++
		}
	}

	fmt.Printf("Percent of students assigned to first choice recitation: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctPrefRec() {
	total := 0
	for _, fn := range t.l {
		if fn.st.rp[0] == fn.rsec.time || fn.st.rp[1] == fn.rsec.time {
			total++
		}
	}

	fmt.Printf("Percent of students assigned to first or second choice recitation: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctBestTut() {
	total := 0
	for _, fn := range t.l {
		if fn.st.tp == fn.tsec.time {
			total++
		}
	}

	fmt.Printf("Percent of students assigned to first choice tutorial: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptTutEnrollment() {
	tuts := map[*section]int{}

	for _, fn := range t.l {
		tut := fn.tsec
		if _, ok := tuts[tut]; !ok {
			tuts[tut] = 1
		} else {
			tuts[tut] += 1
		}
	}

	fmt.Printf("Tutorial enrollment:\n")
	total := 0
	for tut, count := range tuts {
		fmt.Printf("\t%s %s: %d\n", tut.instructor, tut.time, count)
		total += count
	}
	fmt.Printf("Total tutorial enrollment: %d\n", total)
}

func (t *stats) rptSPerT() {
	tuts := map[*section]bool{}

	for _, fn := range t.l {
		tut := fn.tsec
		if _, ok := tuts[tut]; !ok {
			tuts[tut] = true
		}
	}

	fmt.Printf("Students per tutorial: %d\n", len(t.l)/len(tuts)+allowedTutOvf)
}

func (t *stats) prepare(l []*fn) error {
	t.l = l
	return nil
}

func (t *stats) execute() error {
	t.rptPctAssigned()
	t.rptPctBestRec()
	t.rptPctPrefRec()
	t.rptPctBestTut()
	t.rptSPerT()
	t.rptTutEnrollment()
	return nil
}

func (t *stats) export() ([]*fn, error) {
	return t.l, nil
}
