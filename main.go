package main

func main() {
	l := make([]*fn, 0)
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
	if err := exportFns(l); err != nil {
		panic(err)
	}
}
