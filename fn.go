package main

import (
	"fmt"
)

const (
	kid = iota
	rtunion
	tutorial
	src
	sink
)

type fn struct {
	t int

	st   *student
	rsec *section
	tsec *section
}

func (f fn) String() string {
	switch f.t {
	case kid:
		return f.st.String()
	case rtunion:
		return fmt.Sprintf("Union of:\n\t%s\n\t%s\n", f.rsec.String(), f.tsec.String())
	case tutorial:
		return f.tsec.String()
	case src:
		return "Source"
	default:
		return "Sink"
	}
}
