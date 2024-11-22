package p

import (
	"mockery-issue/lib"
)

type T = int
type S = lib.S

//go:generate mockery --quiet --name I
type I interface {
	F(T, S, lib.S)
}
