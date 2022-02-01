package main

import (
	"fmt"
	"testing"
)

func TestSFnGen(t *testing.T) {
	sfns, err := makeStudentFns()
	if err != nil {
		t.Error(err)
	}

	for _, sfn := range sfns {
		if sfn.st.email == "alvvchan@mit.edu" {
			fmt.Println(sfn)
		}
	}
}

func TestRTUGen(t *testing.T) {
	rtus, _ := makeRTUFns()
	for _, rtu := range rtus {
		fmt.Println(rtu)
	}
}

func TestTutGen(t *testing.T) {
	tuts, _ := makeTutFns()
	for _, tut := range tuts {
		fmt.Println(tut)
	}
}
