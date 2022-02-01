package main

import (
	"fmt"
	"testing"
)

func TestLegalRTGen(t *testing.T) {
	lrts := legalRTSectPairs()
	for _, lrt := range lrts {
		fmt.Printf("%s -> %s\n", rbs[lrt[0]], tbs[lrt[1]])
	}
}
