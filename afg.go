package main

import "fmt"

type afg struct {
	fg *fg

	srcfn  *fn
	sinkfn *fn

	sfns  []*fn
	rtfns []*fn
	tfns  []*fn

	res *fgr
}

/* MARK: preparation */

func (a *afg) prepFns(sfns []*fn) error {
	var err error

	a.srcfn = &fn{t: src}
	a.sinkfn = &fn{t: sink}
	a.sfns = sfns

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
	everyFn := append(a.sfns, a.rtfns...)
	everyFn = append(everyFn, a.tfns...)
	everyFn = append(everyFn, a.srcfn, a.sinkfn)

	a.fg = newFg(everyFn)

	for _, sfn := range a.sfns {
		a.fg.addCost(a.srcfn, sfn, 1)
	}

	for _, sfn := range a.sfns {
		for _, rtfn := range a.rtfns {
			if sfn.st.open(rtfn.rsec.time) && sfn.st.open(rtfn.tsec.time) {
				a.fg.addCost(sfn, rtfn, 1)
			}
		}

	}

	for _, rtfn := range a.rtfns {
		for _, tfn := range a.tfns {
			if rtfn.tsec == tfn.tsec {
				a.fg.addCost(rtfn, tfn, int(^uint(0)>>1))
			}
		}
	}

	tutSize := len(a.sfns)/len(a.tfns) + allowedTutOvf
	for _, tfn := range a.tfns {
		a.fg.addCost(tfn, a.sinkfn, tutSize)
	}

	return nil
}

func (a *afg) prepare(sfns []*fn) error {
	err := a.prepFns(sfns)
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
	a.res = a.fg.runMaxFlow(a.srcfn, a.sinkfn)
	if a.res.flow != len(a.sfns) {
		return fmt.Errorf("flow: %d, expected: %d", a.res.flow, len(a.sfns))
	}
	return nil
}

func (a *afg) export() ([]*fn, error) {
	for _, sfn := range a.sfns {
		for _, rtfn := range a.rtfns {
			f := a.res.flowAlong(sfn, rtfn)
			if f > 0 {
				sfn.rsec = rtfn.rsec
				sfn.tsec = rtfn.tsec
			}
		}
	}

	return a.sfns, nil
}
