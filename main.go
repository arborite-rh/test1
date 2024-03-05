package main

import "fmt"

type MyWorker interface {
	Domain1() any
	Domain2() any
}

type worker struct{}

func NewWorker() MyWorker {
	return &worker{}
}

func main() {
	w := NewWorker()
	fmt.Println(w.Domain1())
	fmt.Println(w.Domain2())
}
