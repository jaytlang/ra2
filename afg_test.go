package main

import (
	"fmt"
	"testing"
)

func TestAfg(t *testing.T) {
	afg := &afg{}
	for i := 0; i < 10; i++ {
		afg.prepare()
		fmt.Println(len(afg.fg.capacity), len(afg.fg.adjacency))
		afg.execute()
	}
}
