package p

type T = int

//go:generate mockery --quiet --name I
type I interface{ F(T) }
