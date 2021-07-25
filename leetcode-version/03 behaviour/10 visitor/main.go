package main

import "fmt"

type IVisitor interface {
	Visit()
}

type WeiBoVisitor struct{}

func (w WeiBoVisitor) Visit() {
	fmt.Println("Visit WeiBo")
}

type IQIYIVisitor struct{}

func (i IQIYIVisitor) Visit() {
	fmt.Println("Visit IQiYi")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct{}

func (e Element) Accept(v IVisitor) {
	v.Visit()
}

func main() {
	e := new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}
