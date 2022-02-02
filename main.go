package main

func main() {
	a := afg{}
	a.prepare(nil)
	err := a.execute()
	if err != nil {
		panic(err)
	}
	r, _ := a.export()

	s := asbp{}
	s.prepare(r)
	s.execute()
	r, _ = s.export()

	if err := exportFns(r); err != nil {
		panic(err)
	}
}
