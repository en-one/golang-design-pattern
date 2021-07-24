package main

import "fmt"

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

//==================================
type Machine struct {
	current State
}

func NewMachine() *Machine {
	return &Machine{NewOff()}
}

func (m *Machine) setCurrent(s State) {
	m.current = s
}

func (m *Machine) On() {
	m.current.On(m)
}

func (m *Machine) Off() {
	m.current.Off(m)
}

//==================================
type On struct{}

func NewOn() State {
	return &On{} //On,实现了 on 以及off方法，多态的方式所以继承了state
}

func (o *On) On(m *Machine) {
	fmt.Println("已经开启")
}

func (o *On) Off(m *Machine) {
	fmt.Println("从On的状态到Off")
	m.setCurrent(NewOff())
}

//===================================
type Off struct{}

func NewOff() State {
	return &Off{} //Off,实现了 on 以及off方法，所以继承了state
}

func (o *Off) On(m *Machine) {
	fmt.Println("从OFF的状态到ON")
	m.setCurrent(NewOn())
}

func (o *Off) Off(m *Machine) {
	fmt.Println("已经关闭")
}

//============================
func main() {
	machine := NewMachine() //*main.Off
	machine.Off()           //*main.Off.Off()
	machine.On()            //*main.Off.On();*main.on
	machine.On()            //*main.Off.On()
	machine.Off()           //*main.On.Off();*main.off
}
