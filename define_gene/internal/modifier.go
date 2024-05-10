package internal

//go:generate stringer -type=Dir
type Dir int

const (
	ONC Dir = iota
	ENT
	EXT
)
