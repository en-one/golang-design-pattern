package main

import "fmt"

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value}
}

func (n *Number) Double() {
	n.value = 2 * n.value
}

func (n *Number) Half() {
	n.value = n.value / 2
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}

func (n *Number) ReinstateMemento(memento *Memento) {
	n.value = memento.state
}

//备忘录

type Memento struct {
	state int
}

func NewMemento(value int) *Memento {
	return &Memento{value}
}

func main() {
	n := NewNumber(10)
	n.Double()
	memento := n.CreateMemento() //创建备忘录
	n.Half()                     //修改原址 变一般
	n.ReinstateMemento(memento)  //利用备忘录恢复
	fmt.Println(memento, n)
}
