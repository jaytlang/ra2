package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func makeStudentFns() ([]fn, error) {
	f, err := os.Open(csvFile)

	if err != nil {
		return nil, fmt.Errorf("failed to open config-provided csv %s: %v", csvFile, err)
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read config-provided csv %s: %v", csvFile, err)
	}

	e2fn := make(map[string]fn)
	e2dpe := make(map[string][]string)

	for i, line := range lines {
		if i == 0 {
			continue
		}

		s := &student{
			name:  line[name],
			email: strings.ToLower(line[kerb]) + "@mit.edu",
			avail: make([]st, 0),
			rp:    []st{st(line[rpref1]), st(line[rpref2])},
			tp:    st(line[tpref]),
		}

		avail := strings.Replace(line[ravail], `"`, "", -1)
		availl := strings.Split(avail, ", ")
		for _, a := range availl {
			s.avail = append(s.avail, st(a))
		}

		e2fn[s.email] = fn{
			t:  kid,
			st: s,
		}

		sdpk1 := strings.ToLower(line[dpk1])
		sdpk2 := strings.ToLower(line[dpk2])
		e2dpe[s.email] = make([]string, 0)
		if sdpk1 != "" {
			e2dpe[s.email] = append(e2dpe[s.email], sdpk1)
		}
		if sdpk2 != "" {
			e2dpe[s.email] = append(e2dpe[s.email], sdpk2)
		}
	}

	for em, fn := range e2fn {
		dpfavs := e2dpe[em]
		for _, dpfav := range dpfavs {
			if dpfav == em {
				continue
			}
			fn.st.dpfavs = append(fn.st.dpfavs, e2fn[dpfav].st)
		}
	}

	ret := make([]fn, 0)
	for _, fn := range e2fn {
		ret = append(ret, fn)
	}

	return ret, nil
	/* totally broken lol TODO */
}
