package main

import "fmt"

type Person struct {
	name string
	cmd  Command
}

type Command struct {
	person *Person
	method func()
}

func NewCommand(p *Person, method func()) Command {
	return Command{
		person: p,
		method: method,
	}
}

func (c *Command) Execute() {
	c.method()
}

func NewPerson(name string, cmd Command) Person {
	return Person{
		name: name,
		cmd:  cmd,
	}
}

func (p *Person) Talking() {
	fmt.Println(fmt.Sprintf("%s is Tailking", p.name))
	p.cmd.Execute()

}

func (p *Person) Listen() {
	fmt.Println(fmt.Sprintf("%s is Listening", p.name))
	p.cmd.Execute()
}

func (p *Person) Do() {
	fmt.Println(fmt.Sprintf("%s is Doing", p.name))
}

func main() {
	XiaoZhang := NewPerson("xiaozhang", NewCommand(nil, nil))
	laoZhang := NewPerson("zhang", NewCommand(&XiaoZhang, XiaoZhang.Do))
	laoChen := NewPerson("chen", NewCommand(&laoZhang, laoZhang.Listen))
	laoChen.Talking()
}
