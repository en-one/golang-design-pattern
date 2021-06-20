package main

import "fmt"

//通过组合来代替继承，主要的作用是给原始类添加增强功能
type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s *Square) Draw() string {
	return "this is a square"
}

//实现方式: 继承

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{
		color:  color,
		square: square,
	}
}

func (c *ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}

func main() {
	sq := &Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	fmt.Println(got)
}
