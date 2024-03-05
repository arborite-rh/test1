package main

import "fmt"

type MyWorker interface {
	Domain1(any) any
	Domain2(any) any
}

type worker struct{}

func NewWorker() MyWorker {
	return &worker{}
}

func main() {
	fmt.Println("Hello World")
}
