package main

import "fmt"

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(v int) *Leaf {
	return &Leaf{value: v}
}

func (l *Leaf) Traverse() {
	fmt.Println(l.value)
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}

//怎么说 可能是搞了个递归
func main() {
	containers := make([]Composite, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			containers[i].Add(NewLeaf(i*3 + j))
		}
	}

	for i := 1; i < 4; i++ {
		containers[0].Add(&(containers[i]))
	}

	for i := 0; i < 4; i++ {
		containers[i].Traverse()
		fmt.Printf("\n")
	}

}
