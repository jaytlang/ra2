package main

import (
	"fmt"
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
	str += "\n"
	return str
}

func (s student) open(t st) bool {
	for _, st := range s.avail {
		if st == t {
			return true
		}
	}
	return false
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
