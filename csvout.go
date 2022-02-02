package main

import (
	"encoding/csv"
	"os"
)

func exportFns(fns []*fn) error {
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
			record = append(record, st.email)
		}
		records = append(records, record)
	}

	if err := csvw.WriteAll(records); err != nil {
		return err
	}

	return nil
}
