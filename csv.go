package main

import "strings"

// CSV file format: invariant
const (
	timestamp = iota
	email
	kerb
	name
	ravail
	rpref1
	rpref2
	tavail
	tpref
	dpk1
	dpn1
	dpk2
	dpn2
)

func emailFromKerb(kerb string) string {
	return strings.ToLower(kerb) + "@mit.edu"
}

func csvlName(l []string) string {
	return l[name]
}

func csvlEmail(l []string) string {
	return emailFromKerb(l[kerb])
}

func csvlAvail(l []string) []st {
	availnq := strings.Replace(l[ravail], `"`, "", -1)
	availl := strings.Split(availnq, ", ")
	avail := make([]st, 0)
	for _, a := range availl {
		avail = append(avail, st(a))
	}
	return avail
}

func csvlRp(l []string) []st {
	return []st{st(l[rpref1]), st(l[rpref2])}
}

func csvlTp(l []string) st {
	return st(l[tpref])
}

func csvlDpFav1(l []string) string {
	return emailFromKerb(l[dpk1])
}

func csvlDpFav2(l []string) string {
	return emailFromKerb(l[dpk2])
}
