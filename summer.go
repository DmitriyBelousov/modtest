package modtest

type Summer interface {
	Sum(x, y int) int
}

type summer struct {
}

func (s *summer) Sum(x, y int) int {
	return x + y
}

func New() Summer {
	return &summer{}
}
