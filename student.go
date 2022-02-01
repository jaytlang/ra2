package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type student struct {
	name  string
	email string
	avail []st
	rp    []st
	tp    st

	dpfavs []*student
}

func (s student) String() string {
	str := fmt.Sprintf("%s (%s):\n\tavailable %v\n\twants rs %v\n\twants tut %v\n",
		s.name, s.email, s.avail, s.rp, s.tp)
	str += "\tWants partners: "
	for _, dpfav := range s.dpfavs {
		str += fmt.Sprintf("\n\t\t%s, ", dpfav.email)
	}
	return str
}

func newStudentIso(line []string) *student {
	s := student{
		name:  csvlName(line),
		email: csvlEmail(line),
		avail: csvlAvail(line),
		rp:    csvlRp(line),
		tp:    csvlTp(line),
	}

	return &s
}

func (s *student) addPrefs(e2s map[string]*student, e2dpfavs map[string][]string) {
	for _, dpfav := range e2dpfavs[s.email] {
		if st, ok := e2s[dpfav]; ok {
			s.dpfavs = append(s.dpfavs, st)
		}
	}
}

func newStudentsAll() ([]*student, error) {
	f, err := os.Open(csvFile)

	if err != nil {
		return nil, fmt.Errorf("failed to open config-provided csv %s: %v", csvFile, err)
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read config-provided csv %s: %v", csvFile, err)
	}

	e2s := make(map[string]*student)
	e2dpfavs := make(map[string][]string)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		s := newStudentIso(line)

		e2s[s.email] = s
		e2dpfavs[s.email] = []string{csvlDpFav1(line), csvlDpFav2(line)}
	}

	ret := make([]*student, 0)
	for _, s := range e2s {
		s.addPrefs(e2s, e2dpfavs)
		ret = append(ret, s)
	}

	return ret, nil
}
