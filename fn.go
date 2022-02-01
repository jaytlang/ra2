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

func emailFromKerb(kerb string) string {
	return strings.ToLower(kerb) + "@mit.edu"
}

func makeStudentNP(line []string) student {
	s := student{
		name:  line[name],
		email: emailFromKerb(line[kerb]),
		avail: make([]st, 0),
		rp:    []st{st(line[rpref1]), st(line[rpref2])},
		tp:    st(line[tpref]),
	}

	availnq := strings.Replace(line[ravail], `"`, "", -1)
	availl := strings.Split(availnq, ", ")
	for _, a := range availl {
		s.avail = append(s.avail, st(a))
	}

	return s
}

func addStudentPrefs(s *student, email2fn map[string]fn, email2dpfavs map[string][]string) {
	for _, dpfav := range email2dpfavs[s.email] {
		if fn, ok := email2fn[dpfav]; ok {
			s.dpfavs = append(s.dpfavs, fn.st)
		}
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

	email2fn := make(map[string]fn)
	email2dpfavs := make(map[string][]string)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		s := makeStudentNP(line)

		fn := fn{
			t:  kid,
			st: &s,
		}

		email2fn[s.email] = fn
		email2dpfavs[s.email] = []string{emailFromKerb(line[dpk1]), emailFromKerb(line[dpk2])}
	}

	ret := make([]fn, 0)
	for _, fn := range email2fn {
		addStudentPrefs(fn.st, email2fn, email2dpfavs)
		ret = append(ret, fn)
	}

	return ret, nil
}
