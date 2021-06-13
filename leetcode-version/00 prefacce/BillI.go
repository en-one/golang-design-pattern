package main

type bill interface {
	disCount() float64
}

func measure(b bill) {
	b.disCount()
}
