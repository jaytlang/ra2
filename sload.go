package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func importSfns() ([]*fn, error) {
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

	fns := make([]*fn, 0)

	for _, s := range e2s {
		s.addPrefs(e2s, e2dpfavs)
		fns = append(fns, &fn{
			t:  kid,
			st: s,
		})
	}

	return fns, nil
}

func exportSfns(fns []*fn) error {
	f, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer f.Close()

	records := make([][]string, 0)
	csvw := csv.NewWriter(f)
	defer csvw.Flush()

	hdr := []string{"email", "recitation instructor", "recitation time", "tutorial instructor", "tutorial time", "teammates..."}
	if err := csvw.Write(hdr); err != nil {
		return err
	}

	for _, fn := range fns {
		record := make([]string, 0)
		record = append(record, fn.st.email)

		record = append(record, fn.rsec.instructor)
		record = append(record, string(fn.rsec.time))

		record = append(record, fn.tsec.instructor)
		record = append(record, string(fn.tsec.time))

		for _, st := range fn.team {
			record = append(record, st.st.email)
		}
		records = append(records, record)
	}

	if err := csvw.WriteAll(records); err != nil {
		return err
	}

	return nil
}
