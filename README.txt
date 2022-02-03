==============
     ra2
==============

Next generation recitation assignment tool for 6.033.
Modular, scalable, fast. Designed to be iterated on by
future years.

Features:
- strategies: a unit of computation and an intermediate
	step in the process of mapping students to
	recitations. these share a common interface and
	execute in sequence. additional strategies can
	be added at will to perform logging, optimize
	placements according to various preferences, or
	flat out change the entire core ra2 algorithm.
	sort of like a meta-programming language - the
	full potential of strategies is unexplored.

- afg: a powerful max-flow algorithm that finds =optimal=
	times for students to attend recitation and tutorial.
	100% of students are assigned a compatible recitation,
	and 95% or so are given one of their top two choices.
	non-deterministic: simply re-run to get a different
	mapping if you don't like your current result.

- teams: optimized team placement + recitation swapping achieves
	>50% student satisfaction while still conforming to
	aggressive scheduling constraints + guidelines set forth
	by WRAP.

- speed: the afg binary executes in 0.01s and is largely
	held up by file I/O rather than execution proper.

- server: a sh script to prepare ra2 outputs to be served
	over http. tells your friends / fellow students
	your progress + a copy of the latest outputs from
	all strategies + statistics from the stats strategy
	so that they don't have to keep asking you (:

This tool is written in Go. Download the go distribution
for your platform, and type 'go build' to compile the
ra2 distribution. Then simply execute with ./ra2.
You'll probably need to `mkdir -p data/` so that the
default config works, and make sure that ra2 can pick
up your Google Form outputs.

See config.go for how to use / documentation, or to
change the behavior of ra2, edit the values in config.go
If you're working with a different form, you probably
have to edit additional files (mainly csvin.go) to
ensure everything works properly.

