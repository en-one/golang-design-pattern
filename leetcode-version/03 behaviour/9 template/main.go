package main

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}

type worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *worker {
	return &worker{w}
}

func (w *worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct{}

func (c *Coder) GetUp() {
	fmt.Println("coder GetUp.")
}

func (c *Coder) Work() {
	fmt.Println("coder Coding.")
}

func (c *Coder) Sleep() {
	fmt.Println("coder Sleep.")
}

func main() {
	worker := NewWorker(&Coder{})

	worker.Daily()
}
