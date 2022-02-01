package main

import "fmt"

type afg struct {
	fg *fg

	srcfn  *fn
	sinkfn *fn

	sfns  []*fn
	rtfns []*fn
	tfns  []*fn
}

/* MARK: preparation */

func (a *afg) prepFns() error {
	var err error

	a.srcfn = &fn{t: src}
	a.sinkfn = &fn{t: sink}

	a.sfns, err = makeStudentFns()
	if err != nil {
		return err
	}

	a.rtfns, err = makeRTUFns()
	if err != nil {
		return err
	}

	a.tfns, err = makeTutFns()
	if err != nil {
		return err
	}

	return nil
}

func (a *afg) prepFg() error {
	a.fg = newFg()

	i := 0
	for _, sfn := range a.sfns {
		i++
		a.fg.addEdge(a.srcfn, sfn, 1)
	}

	fmt.Println("source node edges added: ", i)

	i = 0
	for _, sfn := range a.sfns {
		satisfiable := false
		for _, rtfn := range a.rtfns {
			if sfn.st.open(rtfn.rsec.time) && sfn.st.open(rtfn.tsec.time) {
				satisfiable = true
				i++
				a.fg.addEdge(sfn, rtfn, 1)
			}
		}

		if !satisfiable {
			return fmt.Errorf("%s: no satisfiable rtunion", sfn)
		}
	}

	fmt.Println("rtunion node edges added: ", i)

	i = 0
	for _, rtfn := range a.rtfns {
		for _, tfn := range a.tfns {
			if rtfn.tsec == tfn.tsec {
				i++
				a.fg.addEdge(rtfn, tfn, int(^uint(0)>>1))
			}
		}
	}
	fmt.Println("tutorial node edges added: ", i)

	i = 0
	tutSize := len(a.sfns)/len(a.tfns) + 3
	for _, tfn := range a.tfns {
		a.fg.addEdge(tfn, a.sinkfn, tutSize)
		i++
	}

	fmt.Println("sink edges added: ", i)

	return nil
}

func (a *afg) prepare() error {
	err := a.prepFns()
	if err != nil {
		return err
	}

	err = a.prepFg()
	if err != nil {
		return err
	}

	return nil
}

func (a *afg) execute() error {
	fmt.Println("NUMBER OF STUDENTS: ", len(a.sfns))
	flow := a.fg.maximize(a.srcfn, a.sinkfn)
	fmt.Println("FLOW: ", flow)
	return nil
}
