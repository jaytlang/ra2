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
	s.export()
}
