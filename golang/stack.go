package main

type Stack []uint

func (s *Stack) Push(v uint) {
	*s = append(*s, v)
}

func (s *Stack) Pop() uint {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}
