package main

import (
	"fmt"
)

type Mediator interface {
	Communicate(who string)
}

type WildStallion interface {
	SetMediator(mediator Mediator)
}

type Bill struct{ mediator Mediator }

func (b *Bill) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *Bill) Response() {
	fmt.Println("bill:what")
	b.mediator.Communicate("bill")
}

type Ted struct{ mediator Mediator }

func (t *Ted) Talk() {
	fmt.Println("Ted: Bill?")
	t.mediator.Communicate("Ted")
}

func (t *Ted) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

func (t *Ted) Response() {
	fmt.Println("Ted:hahahhahahahahh")
}

//中介者

type ConcreteMediator struct {
	Bill
	Ted
}

func (c *ConcreteMediator) Communicate(who string) {
	if who == "Ted" {
		c.Bill.Response()
	} else {
		c.Ted.Response()
	}
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func main() {
	mediator := NewMediator()
	mediator.Ted.Talk()
}
