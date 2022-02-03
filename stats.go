package main

import (
	"fmt"
	"os"
)

type stats struct {
	l []*fn
	f *os.File
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

	fmt.Fprintf(t.f, "Percent of students assigned to a fitting section: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctBestRec() {
	total := 0
	for _, fn := range t.l {
		if fn.st.rp[0] == fn.rsec.time {
			total++
		}
	}

	fmt.Fprintf(t.f, "Percent of students assigned to first choice recitation: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctPrefRec() {
	total := 0
	for _, fn := range t.l {
		if fn.st.rp[0] == fn.rsec.time || fn.st.rp[1] == fn.rsec.time {
			total++
		}
	}

	fmt.Fprintf(t.f, "Percent of students assigned to first or second choice recitation: %s\n", percent{total, len(t.l)})
}

func (t *stats) rptPctBestTut() {
	total := 0
	for _, fn := range t.l {
		if fn.st.tp == fn.tsec.time {
			total++
		}
	}

	fmt.Fprintf(t.f, "Percent of students assigned to first choice tutorial: %s\n", percent{total, len(t.l)})
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

	fmt.Fprintf(t.f, "Tutorial enrollment:\n")
	total := 0
	for tut, count := range tuts {
		fmt.Fprintf(t.f, "\t%s %s: %d\n", tut.instructor, tut.time, count)
		total += count
	}
}

func (t *stats) rptRecEnrollment() {
	recs := map[*section]int{}

	for _, fn := range t.l {
		rec := fn.rsec
		if _, ok := recs[rec]; !ok {
			recs[rec] = 1
		} else {
			recs[rec] += 1
		}
	}

	fmt.Fprintf(t.f, "Recitation enrollment:\n")
	total := 0
	for rec, count := range recs {
		fmt.Fprintf(t.f, "\t%s %s: %d\n", rec.instructor, rec.time, count)
		total += count
	}

	fmt.Fprintf(t.f, "Total recitation enrollment by time:\n")
	enrollmentByTime := map[st]int{}
	for _, fn := range t.l {
		enrollmentByTime[fn.rsec.time]++
	}
	for tt, count := range enrollmentByTime {
		fmt.Fprintf(t.f, "\t%s: %d\n", tt, count)
	}

}

func (t *stats) rptSPerT() {
	tuts := map[*section]bool{}

	for _, fn := range t.l {
		tut := fn.tsec
		if _, ok := tuts[tut]; !ok {
			tuts[tut] = true
		}
	}

	fmt.Fprintf(t.f, "Students per tutorial: %d\n", len(t.l)/len(tuts)+allowedTutOvf)
}

func (t *stats) rptTeamSatisfaction() {
	pt := 0

	perfs := 0
	pairs := 0
	singles := 0

	for _, fn := range t.l {
		favs := fn.st.dpfavs
		team := fn.team

		wanted := len(fn.st.dpfavs)
		if wanted == 0 {
			continue
		}

		pt++
		got := 0

		for _, fav := range favs {
			for _, tm := range team {
				if fav == tm.st {
					got++
				}
			}
		}

		if got == 2 {
			perfs++
		} else if got == 1 {
			if wanted == 2 {
				pairs++
			} else {
				perfs++
			}
		} else {
			singles++
		}
	}

	fmt.Fprintf(t.f, "Team preferences (percentages calced only with students who wanted teams):\n")
	fmt.Fprintf(t.f, "\tStudents getting all of their team preferences: %s (%d)\n", percent{perfs, pt}, perfs)
	fmt.Fprintf(t.f, "\tStudents getting one of their two team preferences: %s (%d)\n", percent{pairs, pt}, pairs)
	fmt.Fprintf(t.f, "\tStudents getting none of their team preferences: %s (%d)\n", percent{singles, pt}, singles)
	fmt.Fprintf(t.f, "\t[Students with no team preferences: %d]\n", len(t.l)-pt)
}

func (t *stats) rptSubmissions() {
	fmt.Fprintf(t.f, "Number of submissions: %d\n", len(t.l))
}

func (t *stats) rptLargeTeams() {
	lt3 := map[*student]bool{}
	lt4 := map[*student]bool{}
	lt5 := map[*student]bool{}

	for _, fn := range t.l {
		if len(fn.team) == 2 {
			lt3[fn.st] = true
		} else if len(fn.team) == 3 {
			lt4[fn.st] = true
		} else if len(fn.team) == 4 {
			lt5[fn.st] = true
		}
	}

	fmt.Fprintf(t.f, "3-member teams: %d\n", len(lt3)/3)
	fmt.Fprintf(t.f, "4-member teams: %d\n", len(lt4)/4)
	fmt.Fprintf(t.f, "5-member teams: %d\n", len(lt5)/5)
}

func (t *stats) prepare(l []*fn) error {
	t.l = l
	var err error

	t.f, err = os.Create(statFile)
	if err != nil {
		return err
	}
	return nil
}

func (t *stats) execute() error {
	t.rptPctAssigned()
	t.rptPctBestRec()
	t.rptPctPrefRec()
	t.rptPctBestTut()

	t.rptSPerT()
	t.rptTutEnrollment()
	t.rptRecEnrollment()

	t.rptTeamSatisfaction()
	t.rptLargeTeams()

	t.rptSubmissions()
	return nil
}

func (t *stats) export() ([]*fn, error) {
	t.f.Close()
	return t.l, nil
}
