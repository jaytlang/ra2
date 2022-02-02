package main

import "fmt"

func main() {
	a := afg{}
	a.prepare(nil)
	err := a.execute()
	if err != nil {
		panic(err)
	}

	r, _ := a.export()
	for _, sfn := range r {
		fmt.Println(sfn)
	}

	s := asbp{}
	s.prepare(r)
	s.execute()
	s.export()
}
