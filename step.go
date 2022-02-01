package main

type step interface {
	prepare() error
	execute() error
	export() error
}
