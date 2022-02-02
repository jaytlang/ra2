package main

import (
	"testing"
)

func TestAfg(t *testing.T) {
	afg := &afg{}
	for i := 0; i < 10; i++ {
		afg.prepare()
		afg.execute()
	}
}
