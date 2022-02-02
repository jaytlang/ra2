package main

type strategy interface {
	prepare([]*fn) error
	execute() error
	export() ([]*fn, error)
}
