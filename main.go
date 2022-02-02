package main

func main() {
	l, err := importSfns()
	if err != nil {
		panic(err)
	}

	for _, s := range strategies {
		var err error
		if err = s.prepare(l); err != nil {
			panic(err)
		}
		if err = s.execute(); err != nil {
			panic(err)
		}
		l, err = s.export()
		if err != nil {
			panic(err)
		}
	}
	if err := exportSfns(l); err != nil {
		panic(err)
	}
}
