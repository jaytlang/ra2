package main

import (
	"fmt"
	"testing"
)

func TestStudentFns(t *testing.T) {
	ns, err := newStudentsAll()
	if err != nil {
		t.Error(err)
	}
	for _, s := range ns {
		if s.email == "linhvo@mit.edu" {
			fmt.Println(s)
		}
	}
}
