package main

import (
	"fmt"
	"testing"
)

func TestStudentFns(t *testing.T) {
	fns, err := makeStudentFns()
	if err != nil {
		t.Error(err)
	}
	for _, f := range fns {
		if f.st.email == "alvvchan@mit.edu" {
			fmt.Println(f)
		}
	}
	fmt.Println("Total students: ", len(fns))
}
