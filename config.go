package main

// Strategy: steps to run on a given assignment.
// The core building block of the ra2 distribution.
//
// Each 'strategy' attempts to make progress towards
// arriving at a viable student-recitation matching,
// operating on "fns" (full nodes) that contain information
// about either students, tutorials, recitations, or various
// other node types. An fn is effectively a union type in this
// sense, passed around as a common language for all strategies
// to work with.
//
// Each strategy opens by accepting a list of fns - typically,
// these fns represent a set of students being categorized into
// recitations. This property holds for all strategies included
// in the ra2 distribution. Each strategy does some logic on the
// passed list, and then outputs a new list. These actions
// are encapsulated in prepare(), execute(), and export() methods
// respectively (i.e. ra2 calls prepare(), execute(), and export()
// in sequence).
//
// Feel free to add more strategies! The ra2 distribution is
// deliberately modular and is designed to support this seamlessly.
//
// Currently implemented strategies are as follows:
//	- afg: applied flow graph. Takes a list of students and assigns
// 		them to recitations + tutorials. This utilizes a max flow graph
//		satisfy various constraints, such as:
//			- tutorials have a strict maximum capacity and must be saturated
//				as best as possible
//			- recitations must be relatively balanced
//			- students are assigned to one recitation and one tutorial
//		 	- "teaching teams": all students in a tutorial are taught
//				by the same recitation instructor. A given recitation is
//				mapped to valid tutorials its students may attend - this mapping
//				is configurable below.
//
//		afg populates the rsec and tsec fields of each student fn after running.
//		it is extremely fast (thanks to use of a third-party library) and
//		the algorithm architecture suits the problem at hand well. Additionally,
//		afg is non-deterministic, and produces different results each time it is
//	    invoked.
//
//	- asbp: applied symmetric binary matching. Given a list of student fns
//		all assigned to tutorials, this strategy greedily generates teams within
//		each tutorial. The algorithm first greedily assigns students with other
//		students who _mutually_ (hence 'symmetric') prefer to be together. Then,
//		once team preferences within a given tutorial are satisfied, the algorithm
//		randomly assigns all remaining students to teams. Note that asbp by itself
//		tends to perform poorly on sets of tutorials clustered at the same time,
//		as
//			- students who tend to want to work together tend to collaborate
//				on scheduling, and thus submit similar time preferences
//			- if there are many tutorials at similar times, afg will likely split
//				collaborating students into different tutorials by simple chance,
//				which renders asbp largely powerless to create ideal teams
//
//	- topt: combats the shortcomings of asbp by making additional pass over asbp-
//		generated teams. Iterates over all pairs of students within a team - if
//		this pair did not want to be together, topt attempts to find another student
//		that satisfies the below properties:
//		- is available during the same time frame as one of the pair members
//		- would like to be with one of the pair members
//
//		Upon finding a pair that satisfies these properties, topt swaps the pair
//		and updates the schedules + team listings of all students the strategy is
//		processing. In practice, this overcomes the limitations of asbp and increases
//		overall team satisfaction to about 50% - impressive considering that afg
//		is entirely blind to this.
//
//	- stats: prints out various statistics about the current assignment. makes no
//		changes to the assignment, but is useful for validation + for students/staff
//		to see.

var strategies = []strategy{&afg{}, &asbp{}, &topt{}, &stats{}}

var (
	// csvFile + the CSV routines constitute student input data from
	// e.g. a Google form. Currently the RA2 distribution is relatively
	// monolithic and only supports the 6.033 S22 recitation assignment
	// form, exported from Sheets as a .csv, but this is subject to
	// change in the future. Before any strategies are executed, ra2 imports
	// students and their preferences/availabilities from this file.
	csvFile = "./data/s22.csv"

	// outFile is the file in which the final assignment is encoded
	// the output methods (along with inputs, in sload.go) are responsible
	// for transliterating fn data into .csv format. The results of this land
	// here, after all strategies have run.
	outFile = "./data/out.csv"

	// If the statistics strategy is executed, it
	// will pipe its output here rather than polluting
	// stdout.
	statFile = "./data/stats.txt"
)

// Tutorials are required to be balanced by WRAP as best as
// possible. Obviously, the optimal maximum capacity for a tutorial
// is then the number of students divided by the number of tutorials.
// However, this doesn't always work in practice, due to restrictive
// student availabilities causing e.g. afg or similar strategies to
// fail. If this is the case, the maximum capacity is increased
// by a small constant, allowedTutOvf. Increase this number if
// you see failures due to restrictive schedules but try to keep it small.
const allowedTutOvf = 1

// st: string time. parsed straight off the google form.
// These are just convenience constants for csv parsing
// and should be added or deleted based on the availability
// options you have in your form.
type st string

const (
	tr1011 st = "TR 10-11am ET"
	tr1112 st = "TR 11am-12pm ET"
	tr121  st = "TR 12-1pm ET"
	tr12   st = "TR 1-2pm ET"
	tr23   st = "TR 2-3pm ET"

	f12 st = "F 1-2pm ET"
	f23 st = "F 2-3pm ET"
)

// All recitations.
var ars = map[int]*section{
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

// All tutorials.
var ats = map[int]*section{
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

// Teaching teams. r2t is a map representing for each
// recitation instructor, which tutorial instructors
// are teamed with that recitation instructor. Students assigned
// to a recitation instructor r are free to be assigned to
// any tutorial taught by one of r2t[r]. Change these as
// necessary and ensure that instructor names match precisely
// with those in ars and ats.
var r2t = map[string][]string{
	"Karen Sollins":        {"Brianna Williams"},
	"Howard Shrobe":        {"Elizabeth Stevens"},
	"Henry Corrigan-Gibbs": {"Amy Carleton"},
	"Mohammad Alizadeh":    {"Laura McKee", "Keith Clavin"},
	"Larry Rudolph":        {"Thomas Pickering"},
	"John Feser":           {"David Larson"},
	"Michael Cafarella":    {"Jessie Stickgold-Sarah", "Michael Maune"},
	"Adam Belay":           {"Michael Trice"},
}
