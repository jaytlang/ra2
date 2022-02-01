package main

import "fmt"

// TODO: configure via a nice config file or cmdline flags
// Also document what these do
// TODO: make super cool spreadsheet go getter with google apis
// it's late, aint nobody got time for that

var (
	csvFile = "./data/s22.csv"
)

type st string

const (
	tr1011 st = "TR 10-11am ET"
	tr1112 st = "TR 11-12pm ET"
	tr121  st = "TR 12-1pm ET"
	tr12   st = "TR 1-2pm ET"
	tr23   st = "TR 2-3pm ET"

	f12 st = "F 12-1pm ET"
	f23 st = "F 2-3pm ET"
)

var recitationTimes = []st{tr1011, tr1112, tr121, tr12, tr23}
var tutorialTimes = []st{f12, f23}

type section struct {
	isTutorial bool
	time       st
	instructor string
}

func (sec section) String() string {
	if sec.isTutorial {
		return fmt.Sprintf("Tutorial %s: taught by %s\n", sec.time, sec.instructor)
	} else {
		return fmt.Sprintf("Recitation %s: taught by %s\n", sec.time, sec.instructor)
	}
}

var rbs = map[int]section{
	1:  {isTutorial: false, time: tr1011, instructor: "Karen Sollins"},
	2:  {isTutorial: false, time: tr1011, instructor: "Howard Shrobe"},
	3:  {isTutorial: false, time: tr1011, instructor: "Henry Corrigan-Gibbs"},
	4:  {isTutorial: false, time: tr1112, instructor: "Karen Sollins"},
	5:  {isTutorial: false, time: tr1112, instructor: "Howard Shrobe"},
	6:  {isTutorial: false, time: tr1112, instructor: "Henry Corrigan-Gibbs"},
	7:  {isTutorial: false, time: tr121, instructor: "Larry Rudolph"},
	8:  {isTutorial: false, time: tr12, instructor: "Larry Rudolph"},
	9:  {isTutorial: false, time: tr12, instructor: "John Feser"},
	10: {isTutorial: false, time: tr12, instructor: "Michael Cafarella"},
	11: {isTutorial: false, time: tr12, instructor: "Adam Belay"},
	12: {isTutorial: false, time: tr23, instructor: "John Feser"},
	13: {isTutorial: false, time: tr23, instructor: "Michael Cafarella"},
	14: {isTutorial: false, time: tr23, instructor: "Adam Belay"},
	15: {isTutorial: false, time: tr1112, instructor: "Mohammad Alizadeh"},
	16: {isTutorial: false, time: tr121, instructor: "Mohammad Alizadeh"},
}

var tbs = map[int]section{
	2:  {isTutorial: true, time: f12, instructor: "Laura McKee"},
	3:  {isTutorial: true, time: f12, instructor: "Amy Carleton"},
	4:  {isTutorial: true, time: f23, instructor: "Keith Clavin"},
	5:  {isTutorial: true, time: f23, instructor: "Amy Carleton"},
	6:  {isTutorial: true, time: f12, instructor: "David Larson"},
	7:  {isTutorial: true, time: f23, instructor: "Michael Trice"},
	9:  {isTutorial: true, time: f23, instructor: "Elizabeth Stevens"},
	10: {isTutorial: true, time: f12, instructor: "Elizabeth Stevens"},
	11: {isTutorial: true, time: f23, instructor: "David Larson"},
	12: {isTutorial: true, time: f23, instructor: "Michael Maune"},
	13: {isTutorial: true, time: f12, instructor: "Jessie Stickgold-Sarah"},
	14: {isTutorial: true, time: f23, instructor: "Thomas Pickering"},
	15: {isTutorial: true, time: f23, instructor: "Brianna Williams"},
	16: {isTutorial: true, time: f12, instructor: "Michael Trice"},
	17: {isTutorial: true, time: f12, instructor: "Thomas Pickering"},
	18: {isTutorial: true, time: f12, instructor: "Brianna Williams"},
}

var r2t = map[string][]string{
	"Karen Sollins":        {"Brianna Williams"},
	"Howie Shrobe":         {"Elizabeth Stevens"},
	"Henry Corrigan-Gibbs": {"Amy Carleton"},
	"Mohammad Alizadeh":    {"Laura McKee", "Keith Clavin"},
	"Larry Rudolph":        {"Thomas Pickering"},
	"John Feser":           {"David Larson"},
	"Michael Cafarella":    {"Jessie Stickgold-Sarah", "Michael Maune"},
	"Adam Belay":           {"Michael Trice"},
}

func findSectNs(inst string, rec bool) []int {
	m := tbs
	if rec {
		m = rbs
	}

	sns := []int{}
	for sn, s := range m {
		if s.instructor == inst {
			sns = append(sns, sn)
		}
	}

	return sns
}

func legalRTSectPairs() [][2]int {
	pairs := make([][2]int, 0)

	for ri, tis := range r2t {
		for _, ti := range tis {
			rsns := findSectNs(ri, true)
			tsns := findSectNs(ti, false)
			for _, rsn := range rsns {
				for _, tsn := range tsns {
					pairs = append(pairs, [2]int{rsn, tsn})
				}
			}
		}
	}

	return pairs
}
