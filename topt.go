package main

type topt struct {
	l []*fn
}

func (t *topt) mfvd(a *student, b *student) bool {
	w1 := false
	w2 := false

	for _, dpfav := range a.dpfavs {
		if dpfav == b {
			w1 = true
		}
	}

	for _, dpfav := range b.dpfavs {
		if dpfav == a {
			w2 = true
		}
	}

	return w1 && w2

}

func (t *topt) teammates(a *fn, b *fn) bool {
	for _, tm := range a.team {
		if tm == b {
			return true
		}
	}
	return false
}

func (t *topt) canSwap(a *fn, b *fn) bool {
	if a.st.open(b.rsec.time) && a.st.open(b.tsec.time) {
		if b.st.open(a.rsec.time) && b.st.open(a.tsec.time) {
			return true
		}
	}
	return false
}

func (t *topt) stb(a *fn, b *fn) {
	a.team, b.team = b.team, a.team
	for _, fn := range t.l {
		if fn == a || fn == b {
			continue
		}
		for i, tm := range fn.team {
			if tm == a {
				fn.team[i] = b
			} else if tm == b {
				fn.team[i] = a
			}
		}
	}
}

func (t *topt) prepare(l []*fn) error {
	t.l = l
	return nil
}

func (t *topt) execute() error {
	swpd := make(map[*fn]bool)

	for _, fn := range t.l {
		swpd[fn] = true
		for i, tm := range fn.team {
			tm = fn.team[i]

			if !t.mfvd(fn.st, tm.st) {

				for _, rplc := range t.l {
					if t.teammates(fn, rplc) || fn == rplc {
						continue
					} else if swpd[rplc] {
						continue
					}

					if t.mfvd(fn.st, rplc.st) && t.canSwap(tm, rplc) {
						tm.rsec, rplc.rsec = rplc.rsec, tm.rsec
						tm.tsec, rplc.tsec = rplc.tsec, tm.tsec
						t.stb(tm, rplc)

						swpd[rplc] = true
						swpd[tm] = true
					}
				}
			}
		}
	}

	return nil
}

func (t *topt) export() ([]*fn, error) {
	return t.l, nil
}
