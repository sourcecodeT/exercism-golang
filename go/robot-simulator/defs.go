package robot

import "sync"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
	sync.Mutex
}

type Dir int

const (
	E = iota
	W
	S
	N
)

// additional definitions used in step 2

type Command byte
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}
